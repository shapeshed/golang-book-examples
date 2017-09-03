package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	needle := "(?i)chocolate"
	haystack := "Chocolate is my favorite!"
	match, err := regexp.MatchString(needle, haystack)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(match)
}
