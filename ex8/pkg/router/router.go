package router

import (
	"github.com/alexvux/dwarves-go23/ex8/pkg/handler"
	mw "github.com/alexvux/dwarves-go23/ex8/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	addProductRoutes(r)
	addCartRoutes(r)
	return r
}

func addProductRoutes(r *gin.Engine) {
	productRoutes := r.Group("/products")
	productRoutes.Use(mw.BasicAuth())

	productRoutes.GET("/", handler.GetAllProducts)
	productRoutes.POST("/", handler.AddProduct)
	productRoutes.PUT("/:id", handler.UpdateProduct)
	productRoutes.DELETE("/:id", handler.DeleteProduct)

}

func addCartRoutes(r *gin.Engine) {
	cartRoutes := r.Group("/cart")
	cartRoutes.POST("/", handler.AddItemToCart)
	cartRoutes.DELETE("/:id", handler.DeleteItemFromCart)
	cartRoutes.POST("/checkout", handler.Checkout)
}
