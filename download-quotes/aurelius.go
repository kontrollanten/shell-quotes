package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
)

var (
	contentRegexp *regexp.Regexp = regexp.MustCompile("“(.+?)”")
)

func (q *Quote) String() string {
	return fmt.Sprintf("%s ― %s\n", q.Text, q.Author)
}

func DownloadAurelius() {
	var amount int = 100

	var quotes []Quote

	c := colly.NewCollector(
		colly.AllowedDomains("www.goodreads.com"),
	)

	c.OnHTML(".quoteDetails", func(e *colly.HTMLElement) {
		res := contentRegexp.FindAllStringSubmatch(e.ChildText("div.quoteText"), -1)

		if len(res) < 1 {
			return
		}

		if len(res[0]) < 1 {
			return
		}

		quotes = append(quotes, Quote{
			Text: res[0][0],
			Author: strings.ReplaceAll(
				strings.ReplaceAll(e.ChildText(".authorOrTitle"), "\n", " "),
				"  ",
				"",
			),
		})

		fmt.Print(".")
	})

	c.OnHTML(".next_page", func(e *colly.HTMLElement) {
		if len(quotes) < amount {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.Visit("https://www.goodreads.com/author/quotes/17212.Marcus_Aurelius")

	fmt.Printf("Scrapped %d quotes.\n\n", len(quotes))

	WriteToFile(quotes, "aurelius.txt")
}
