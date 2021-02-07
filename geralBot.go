package main

import (
	"github.com/gocolly/colly"
)

func geralGetter() string {

	c := colly.NewCollector()

	var geral string

	c.OnHTML("h1", func(h *colly.HTMLElement) {
		geral = h.Text
	})
	c.Visit("https://en.wikipedia.org/wiki/Special:Random")

	return geral
}
