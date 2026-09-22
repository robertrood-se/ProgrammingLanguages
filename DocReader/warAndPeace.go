package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var wapUrl = "https://www.gutenberg.org/cache/epub/2600/pg2600.txt"

func main() {
	// load the text file from the webpage
	warAndPeace, urlErr := http.Get(wapUrl)
	if urlErr != nil {
		log.Fatal(urlErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(warAndPeace.Body)
	if warAndPeace.StatusCode != http.StatusOK {
		log.Fatalf("Bad status: %s", warAndPeace.Status)
	}
	body, err := io.ReadAll(warAndPeace.Body)

	// divide up the text into a slice
	words := strings.Fields(string(body))
	// create a map of the word and it's counts
	countsOfWords := make(map[string]int)

	// Now count all the instances of the different words in the body of the webpage
	for _, word := range words {
		// remove any punctuation marks
		word = strings.ReplaceAll(word, "-", " ")
		word = strings.ToLower(word)
		word = strings.TrimRight(word, ".,!?;:\\'()[]{}")
		word = strings.TrimLeft(word, ".,!?;:\\'()[]{}")

		countsOfWords[word] = countsOfWords[word] + 1
	}

	//// Now that we've stripped unnecessary characters out of the words, remove any duplicates
	//seen := make(map[string]struct{})
	//uniqueWords := make([]string, 0)
	//for _, word := range words {
	//	// If the word is not in the map, it's unique
	//	if _, exists := seen[word]; !exists {
	//		seen[word] = struct{}{}
	//		uniqueWords = append(uniqueWords, word)
	//	}
	//}

	// Lastly count the

	// Now print the counts
	for word, count := range countsOfWords {
		fmt.Printf("%s: %d\n", word, count)
	}

	if err != nil {
		return
	}
}
