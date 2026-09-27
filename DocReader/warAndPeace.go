package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"unicode"
)

var wapUrl = "https://www.gutenberg.org/cache/epub/2600/pg2600.txt"

func main() {
	// load the text file from the webpage - Didn't use AI for this part, just looked it up
	warAndPeace, urlErr := http.Get(wapUrl)
	if urlErr != nil {
		log.Fatal(urlErr)
	}
	defer func(BodyText io.ReadCloser) {
		err := BodyText.Close()
		if err != nil {

		}
	}(warAndPeace.Body)
	if warAndPeace.StatusCode != http.StatusOK {
		log.Fatalf("Bad status: %s", warAndPeace.Status)
	}
	rawBody, err := io.ReadAll(warAndPeace.Body)
	// L31-37 are from an AI evaluation of why not all of the special characters were getting removed
	// from the final word list.
	// An interesting point, after I got FieldsFunc working properly using IsLetter()
	// Gemini pointed out to me that L35 & 36 were redundant, but I had to say "thanks" first.
	// clean the body first, maybe that will get rid of the invalid chars
	// cleanedBody := strings.TrimRight(string(rawBody), ".,!?;:'\"_()[]{}")
	// cleanedBody = strings.TrimLeft(string(rawBody), ".,!?;:'\"_()[]{}")
	// divide up the text into a slice
	words := strings.FieldsFunc(string(rawBody), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	// countOfWords is simply a map to keep the list of words and the number of instances of the
	// specific words
	countsOfWords := make(map[string]int)

	// Now count all the instances of the different words in the body of the webpage
	for _, word := range words {
		// remove any punctuation marks and set it to lowercase
		word = strings.ReplaceAll(word, "-", " ")
		word = strings.ToLower(word)

		countsOfWords[word] = countsOfWords[word] + 1
	}

	for word, count := range countsOfWords {
		fmt.Printf("%s: %d\n", word, count)
	}

	if err != nil {
		return
	}
}
