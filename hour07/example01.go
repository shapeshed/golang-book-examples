package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}
	fmt.Println(m.Name, m.Rating)
}
