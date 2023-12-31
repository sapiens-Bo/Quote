package main

import (
	"QuoteDay/data"
	"QuoteDay/display"
)

func main() {
	quote := data.GetData()
	// fmt.Println(quote.Content)
	// fmt.Println(quote.Author)

	display.Display(quote.Content, quote.Author)
}
