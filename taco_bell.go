package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// sendAlert will be responsible for all alerts that are sent after detecting a change
func sendAlert(additions []string, removals []string) {
	if len(additions) != 0 {
		fmt.Println("NEW ITEMS!")
		for _, item := range additions {
			fmt.Println(item)
		}
	}
	if len(removals) != 0 {
		fmt.Println("REMOVED ITEMS!")
		for _, item := range removals {
			fmt.Println(item)
		}
	}
}

// source:
// https://stackoverflow.com/questions/15323767/does-go-have-if-x-in-construct-similar-to-python
// Checks if a string exists within a slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// checkSite is responsible for querying the Taco Bell "new item" page and returning
// an alert if it detects any additions to the lineup.
func checkSite(items *[]string, firstRun bool, debug bool) {
	// Check all product-name divs
	c := colly.NewCollector()
	newItems := make([]string, 0)
	c.OnHTML("div.product-name", func(e *colly.HTMLElement) {
		if firstRun {
			*items = append(*items, strings.Trim(e.Text, " \t\n"))
		}
		{
			newItems = append(newItems, strings.Trim(e.Text, " \t\n"))
		}
	})
	c.Visit("https://www.tacobell.com/food/new")

	// Debug stuff
	if debug && !firstRun {
		newItems = newItems[:len(newItems)-1]
		newItems = append(newItems, "Buffalo Chicken Taco")
	}

	// Check for changes in the page
	additions := make([]string, 0)
	removals := make([]string, 0)
	if !firstRun {
		fmt.Println("Checking for new items...")
		for _, item := range newItems {
			if !stringInSlice(item, *items) {
				additions = append(additions, item)
			}
		}
		for _, item := range *items {
			if !stringInSlice(item, newItems) {
				removals = append(removals, item)
			}
		}
		// Changes were made, send alert
		if len(additions) != 0 || len(removals) != 0 {
			sendAlert(additions, removals)
			*items = newItems
		}
	}
}

func main() {
	items := make([]string, 0)
	checkSite(&items, true, false)

	for {
		checkSite(&items, false, false)
		fmt.Println("Waiting half an hour.")
		time.Sleep(time.Hour / 2)
	}

	//checkSite(&items, false, false)
	//checkSite(&items, false, true)
	//checkSite(&items, false, true)
}
