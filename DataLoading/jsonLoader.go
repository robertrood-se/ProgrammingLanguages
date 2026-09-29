package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type OpenLibraryResponse struct {
	NumFound      int   `json:"num_found"`
	Start         int   `json:"start"`
	NumFoundExact bool  `json:"num_found_exact"`
	Docs          []Doc `json:"docs"`
}

type Doc struct {
	Key                   string   `json:"key"`
	Type                  string   `json:"type"`
	Title                 string   `json:"title"`
	AuthorName            []string `json:"author_name"`
	AuthorKey             []string `json:"author_key"`
	AuthorAlternativeName []string `json:"author_alternative_name"`
	FirstPublishYear      int      `json:"first_publish_year"`
	PublishYear           []int    `json:"publish_year"`
	EditionCount          int      `json:"edition_count"`
	Isbn                  []string `json:"isbn"`
	Language              []string `json:"language"`
	PublishPlace          []string `json:"publish_place"`
	Publisher             []string `json:"publisher"`
	Subject               []string `json:"subject"`
	Person                []string `json:"person"`
	Place                 []string `json:"place"`
	Time                  []string `json:"time"`
	EbookAccess           string   `json:"ebook_access"`
	HasFulltext           bool     `json:"has_fulltext"`
	PublicScanB           bool     `json:"public_scan_b"`
	CoverEditionKey       string   `json:"cover_edition_key"`
	CoverI                int      `json:"cover_i"`
}

func main() {
	url := "https://openlibrary.org/search.json?author=zeihan"

	// Create a custom HTTP client with a reasonable timeout
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	// 1. Fetch the JSON data
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// Purposefully not handling this error
		}
	}(resp.Body)

	// Check that we got a successful status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	// 2. Decode the JSON response
	var result OpenLibraryResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	// 3. Use the parsed data
	fmt.Printf("Total books found: %d\n\n", result.NumFound)

	for i, doc := range result.Docs {
		// Stop after printing the first 5 results to keep the output readable
		if i >= 5 {
			break
		}

		// Handle cases where author might be missing
		author := "Unknown"
		if len(doc.AuthorName) > 0 {
			author = doc.AuthorName[0]
		}

		fmt.Printf("Title: %s\n", doc.Title)
		fmt.Printf("Author: %s\n", author)
		fmt.Printf("First Published: %d\n", doc.FirstPublishYear)
		fmt.Printf("Editions: %d\n", doc.EditionCount)
		fmt.Println("--------------------------------------------------")
	}
}
