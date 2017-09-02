package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "Hello World"
	err := ioutil.WriteFile("example03.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
