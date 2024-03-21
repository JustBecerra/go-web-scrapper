package main

import "github.com/gocolly/colly"

type Industry struct {
	Url, Image, Name string 
}

func main() {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	c.Visit("https://brightdata.com/")
	// var industries []Industry
	c.OnHTML(".product_cards", func(e *colly.HTMLElement){
		
	})
}