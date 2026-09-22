package main

import (
	"fmt"
	"os"
	"strings"
)

func getFile(filename string) []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	bigString := string(contents)
	allLines := strings.Split(bigString, "\n")
	return allLines
}

func main() {
	// Load a file and only print out every other line
	file := getFile("DocReader/silly.txt")
	for i, ln := range file {
		if i%2 != 0 {
			fmt.Println(ln)
		}
	}
}
