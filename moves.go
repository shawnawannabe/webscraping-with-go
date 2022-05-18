package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	Price string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("books.toscrape.com"),
	)

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		fmt.Println(book.Title, book.Price)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://books.toscrape.com/")
}

// func test() {
// 	fmt.Println("Done")
// }
