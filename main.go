package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	fName := "minerals.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	c := colly.NewCollector()

	// Define the URL you want to scrape
	url := "https://de.wikipedia.org/wiki/Liste_mineralischer_Schmuck-_und_Edelsteine"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.OnHTML("html", func(t *colly.HTMLElement) {
		t.ForEach("table", func(_ int, e *colly.HTMLElement) {
			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				m := Mineral{
					Name:          el.ChildText("td:nth-child(1)"),
					Type:          el.ChildText("td:nth-child(2)"),
					Hardness:      el.ChildText("td:nth-child(3)"),
					Density:       el.ChildText("td:nth-child(4)"),
					Crystalsystem: el.ChildText("td:nth-child(5)"),
				}
				fmt.Printf("%+v\n", m)
			})
		})

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Visit the URL and start scraping
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
