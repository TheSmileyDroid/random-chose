package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func themesGetter() map[int]string {

	c := colly.NewCollector()

	var themes map[int]string

	themes = make(map[int]string)

	c.OnHTML(".CategoryTreeSection", func(h *colly.HTMLElement) {
		text := strings.Replace(strings.Split(h.Text, "(")[0], "►  ", "", -1)
		themes[h.Index] = text
	})
	c.Visit("https://en.wikipedia.org/wiki/Category:Video_games_by_theme")

	return themes
}
