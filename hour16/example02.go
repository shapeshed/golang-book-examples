package main

import (
	"errors"
	"log"
)

func main() {
	var errFatal = errors.New("We only just started and we are crashing")
	log.Fatal(errFatal)
}
