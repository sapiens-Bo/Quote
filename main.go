package main

import (
	"QuoteDay/data"
	"QuoteDay/display"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Category string `short:"c" long:"category" default:"Famous Quotes" description:"Category quotes. Default: Famous Quotes."`
}

func main() {
	flags.Parse(&opts)

	quote := data.GetData(opts.Category)
	// fmt.Println(quote.Author)

	display.Display(quote.Content, quote.Author)
}
