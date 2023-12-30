package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Quote struct {
	ID           string   `json:"_id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

func GetData() Quote {
	url := "https://api.quotable.io/random"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error with %s\n", url)
		os.Exit(1)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v, Status code: %v", err.Error(), resp.StatusCode)
	}

	var quote Quote

	err = json.Unmarshal(data, &quote)
	return quote
}
