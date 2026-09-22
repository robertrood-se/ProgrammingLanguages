package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var wapUrl = "https://www.gutenberg.org/cache/epub/2600/pg2600.txt"

func main() {
	// load the text file from the webpage
	warAndPeace, urlErr := http.Get(wapUrl)
	if urlErr != nil {
		log.Fatal(urlErr)
	}
	body, err := io.ReadAll(warAndPeace.Body)
	fmt.Println(body)
	err = warAndPeace.Body.Close()
	if err != nil {
		return
	}
}
