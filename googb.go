package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type googleBooksMetaData struct {
	TotalItems int
	Items      []*item
}

type item struct {
	volinfo `json:"volumeInfo"`
}

type volinfo struct {
	Title         string
	Subtitle      string
	Authors       []string
	Publisher     string
	PublishedDate string
	Description   string
	images        `json:"imageLinks"`
}

type images struct {
	Image string `json:"thumbnail"`
}

var (
	googleBooksAPI    = "https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s"
	googleBooksAPIKey = os.Getenv("GOOGLE_BOOKS_API_KEY")
)

// ScrapeGoogleBooks returns a book struct containing data sourced from the
// Google books API
func ScrapeGoogleBooks(isbn string) (*Book, error) {

	url := fmt.Sprintf(googleBooksAPI, isbn, googleBooksAPIKey)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var result googleBooksMetaData

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.TotalItems == 0 {
		return nil, errors.New("request returned empty meta data")
	}

	d := result.Items[0]
	book := Book{
		d.Title, d.Subtitle, d.Authors, d.Publisher, d.PublishedDate,
		d.Description, d.Image,
	}
	return &book, nil

}
