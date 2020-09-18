package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

type Quote struct {
	Author string
	Text   string
}

func main() {
	targetPath := os.Args[1]

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

		writeToFile(quotes, targetPath)
	})

	c.Visit("http://quotes.cat-v.org/")
}

func writeToFile(quotes []Quote, targetPath string) {
	fmt.Println("targetPath:")
	fmt.Println(targetPath)
	f, e := os.Create(fmt.Sprintf("%s/%s", targetPath, "quotes.txt"))

	if e != nil {
		panic(e)
	}

	defer f.Close()

	for _, q := range quotes {
		row := fmt.Sprintf("%s,%s\n", q.Text, q.Author)
		_, e2 := f.WriteString(row)

		if e2 != nil {
			panic(e2)
		}
	}
}
