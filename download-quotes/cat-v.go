package main

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

func DownloadCatV() {
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.cat-v.org"),
	)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		quotes := []Quote{}
		q := Quote{}

		e.ForEach("*", func(i int, c *colly.HTMLElement) {
			if c.Name == "hr" && q.Text != "" {
				quotes = append(quotes, q)
				q = Quote{}
				return
			}

			if q.Text == "" {
				q.Text = strings.ReplaceAll(strings.TrimSpace(c.Text), "\n", " ")
				return
			}

			q.Author = strings.TrimSpace(strings.TrimSuffix(c.Text, "\n"))
		})

		WriteToFile(quotes, "cat-v.txt")
	})

	c.Visit("http://quotes.cat-v.org/")
}
