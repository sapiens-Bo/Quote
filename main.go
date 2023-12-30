package main

import (
	"QuoteDay/data"
	"fmt"
)

func main() {
	quote := data.GetData()
	fmt.Println(quote.Content)
	fmt.Println(quote.Author)
}
