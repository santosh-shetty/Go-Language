package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type product struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Image string `json:"image"`
}

var products []product

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains:
		colly.AllowedDomains(
			"www.rasmmat.com",
			"rasmmat.com",
			"https://rasmmat.com",
		),
	)

	// Define a callback function that will be called for each page
	c.OnHTML(".pagination .page-item a.page-link", func(e *colly.HTMLElement) {
		// Find the link to the next page
		nextPage := e.Attr("href")
		rel := e.Attr("rel") // if rel="prev" and rel="next" then leave this link
		if nextPage != "" && rel == "" {
			// Visit the next page
			e.Request.Visit(nextPage)
			fmt.Println(rel)
			fmt.Println(nextPage)
		}

	})

	// On every a element which has href attribute call callback
	c.OnHTML("div.aiz-card-box", func(h *colly.HTMLElement) {
		name := h.ChildText("a.text-reset")
		price := h.ChildText(".text-left .fs-15 span.text-primary")
		image := h.ChildAttr("img.img-fit", "data-src")

		products = append(products, product{Name: name,
			Price: price,
			Image: image})

	})
	// Start scraping
	c.Visit("https://rasmmat.com/category/painting-0wppz")

	response, err := json.Marshal(products)
	if err != nil {
		fmt.Println("Error During Marshal products :", err)
	}
	// fmt.Println(string(response))
	os.WriteFile("products.json", response, 755)
}
