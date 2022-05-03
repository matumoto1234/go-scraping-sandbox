package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("submission-code", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://atcoder.jp/contests/abc248/submissions/31042482")
}
