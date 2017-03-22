package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Book describes a book object
type Book struct {
	Title         string
	Subtitle      string
	Authors       []string
	Publisher     string
	PublishedDate string
	Description   string
	Image         string
}

func main() {

	isbn := "9781781688458"

	data, err := ScrapeGoogleBooks(isbn)

	if err != nil {
		log.Fatal(err)
	}

	json, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))

	// isbn := "1469769166"
	// ScrapeOpenLibrary(isbn)

}
