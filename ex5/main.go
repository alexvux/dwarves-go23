package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type Comic struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Image       string `json:"image"`
	LastChapter string `json:"last_chapter"`
	Status      string `json:"status"`
	Views       string `json:"views"`
	Subscribe   string `json:"subcribe"`
}

func main() {
	// remove data file (if existed) before scraping
	filename := "data.json"
	if _, err := os.Stat(filename); err == nil {
		if err := os.Remove(filename); err != nil {
			log.Fatalln("Error on removing data file:", err)
		}
		fmt.Println("Removing old data file before scraping")
	}

	url := "https://truyenqqq.vn/the-loai/action-26.html"
	domains := []string{"https://truyenqqq.vn", "truyenqqq.vn"}
	comics := []Comic{}

	c := colly.NewCollector(colly.AllowedDomains(domains...))
	c.SetRequestTimeout(120 * time.Second)
	extensions.RandomUserAgent(c)

	// set up callbacks
	c.OnHTML("ul.list_grid > li", func(h *colly.HTMLElement) {
		comic := Comic{}
		comic.Name = h.ChildText("div.book_info > div.book_name.qtip > h3 > a")
		comic.Url = h.ChildAttr("div.book_avatar > a", "href")
		comic.Image = h.ChildAttr("div.book_avatar > a > img", "src")
		comic.LastChapter = h.ChildText("div.book_info > div.last_chapter > a")
		comic.Status = h.ChildText("div.book_info > div.more-info > p:nth-child(2)")
		comic.Views = h.ChildText("div.book_info > div.more-info > p:nth-child(3)")
		comic.Subscribe = h.ChildText("div.book_info > div.more-info > p:nth-child(4)")
		comics = append(comics, comic)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Response", r, "got this error:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping", r.Request.URL)
		js, err := json.MarshalIndent(comics, "", "\t")
		if err != nil {
			log.Fatalln(err)
		}

		if err := os.WriteFile(filename, js, 0644); err == nil {
			fmt.Println("Data written to file successfully")
		}
	})

	c.Visit(url)
}
