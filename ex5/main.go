package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Manga struct {
	Name     string
	Chapters int
	Views    int
	Likes    int
	Img      string
	Url      string
}

func main() {
	// remove data file (if existed) before scraping
	filename := "fantasy-manga.json"
	if _, err := os.Stat(filename); err == nil {
		if err := os.Remove(filename); err != nil {
			log.Fatalln("Error on removing data file:", err)
		}
		fmt.Println("Removing old data file before scraping...")
	}

	url := "https://thienhatruyen.com/the-loai/fantasy"
	domains := []string{"https://thienhatruyen.com", "thienhatruyen.com"}
	mangalist := []Manga{}

	c := colly.NewCollector(colly.AllowedDomains(domains...))
	c.SetRequestTimeout(120 * time.Second)
	extensions.RandomUserAgent(c)

	// set up callbacks
	// todo: scrape
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Get a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Got this error:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		js, err := json.MarshalIndent(mangalist, "", "\t")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Write data to file")
		if err := os.WriteFile(filename, js, 0644); err == nil {
			fmt.Println("Data written to file successfully")
		}
	})

	c.Visit(url)
}
