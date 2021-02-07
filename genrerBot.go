package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func genrerGetter() map[int]string {

	c := colly.NewCollector()

	var genres map[int]string

	genres = make(map[int]string)

	c.OnHTML("h3", func(h *colly.HTMLElement) {
		if strings.HasSuffix(h.Text, "[edit]") && h.Text != "Other related topics[edit]" {
			text := strings.Replace(h.Text, "[edit]", "", -1)
			genres[h.Index] = text
		}
	})
	c.Visit("https://en.wikipedia.org/wiki/List_of_video_game_genres")

	return genres
}
