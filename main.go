package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("Hello World")

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce/")
}
