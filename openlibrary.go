package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type openLibraryMataData struct {
	isbndata `json:"ISBN"`
}

type isbndata struct {
	Title         string
	Subtitle      string
	NumberOfPages string `json:"number_of_pages"`
	Publishers    []string
	Authors       []string
	PublishDate   string
	Subjects      []*subjects
}

type subjects struct {
	name string
}

var (
	openLibraryAPI = "http://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json"
)

// ScrapeOpenLibrary returns a book struct containing data sourced from the
// Open Library API
func ScrapeOpenLibrary(isbn string) (*Book, error) {

	url := fmt.Sprintf(openLibraryAPI, isbn)

	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var result openLibraryMataData

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	fmt.Println(&result)

	b := Book{}
	return &b, nil
}
