package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type titulares struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ultimahora.com"),
	)
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scrapping: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
		if r.StatusCode == 200 {
			c.OnHTML("body", func(e *colly.HTMLElement) {
				e.ForEach(".article-title", func(num int, el *colly.HTMLElement) {
					titulares := titulares{
						Title: el.ChildText("a"),
						Link:  el.ChildAttr("a", "href"),
					}
					fmt.Println(titulares)
				})
			})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "nError:", err)
	})

	c.Visit(
		"https://www.ultimahora.com/",
	)
}
