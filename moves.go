package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	Price string
}

// https://stackoverflow.com/questions/1517582/what-is-the-difference-between-statically-typed-and-dynamically-typed-languages
// https://gobyexample.com/structs

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

	// c.OnHTML(".next > a", func(e *colly.HTMLElement) {
	// 	nextPage := e.Request.AbsoluteURL(e.Attr("href"))
	// 	c.Visit(nextPage)
	// })

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://books.toscrape.com/")

	file, err := os.Create("export.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}

// func test() {
// 	fmt.Println("Done")
// }
