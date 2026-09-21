package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type classInfo struct {
	class      string
	instructor string
	credits    int
}

func main() {
	var classes = make([]classInfo, 0)
	fmt.Print("Hello! ")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("can you tell me a class that you are taking? ")
		name, err := reader.ReadString('\n')
		fmt.Print("Who is the instructor? ")
		prof, err := reader.ReadString('\n')
		fmt.Print("How many credits is it? ")
		cred, err := reader.ReadString('\n')
		var credits int
		creditsAsStr := strings.TrimSpace(cred)
		credits, err = strconv.Atoi(creditsAsStr)
		if err != nil {
			panic(err)
		}
		newClass := classInfo{name, prof, credits}
		classes = append(classes, newClass)
		// check to see if they want to continue
		fmt.Print("Do you have another class to enter? ")
		anotherClass, err := reader.ReadString('\n')
		anotherClass = strings.TrimSpace(anotherClass)
		if anotherClass == "no" || anotherClass == "n" {
			break
		}
	}
	for i, class := range classes {
		fmt.Println(class, i)
	} //classes = append(classes, name)
}
