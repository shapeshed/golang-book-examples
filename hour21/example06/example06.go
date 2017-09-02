package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("./deleteme.txt")
	if err != nil {
		log.Fatal(err)
	}
}
