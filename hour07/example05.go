package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := Movie{
		Name:   "Metropolis",
		Rating: 0.99,
	}
	fmt.Printf("%+v\n", m)
}
