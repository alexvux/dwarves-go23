package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexvux/dwarves-go23/ex8/pkg/repo"
	"github.com/alexvux/dwarves-go23/ex8/pkg/router"
)

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=ex8_local port=5432 sslmode=disable"
	if err := repo.InitDB(dsn); err != nil {
		log.Fatal(err)
	}

	router := router.SetupRouter()
	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	// the context is used to inform the server it has 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
