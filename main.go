package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Find("#submission-code").Text())
	})

	c.Visit("https://atcoder.jp/contests/abc248/submissions/31042482")
}
