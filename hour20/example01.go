package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	hobbies := []string{"Cycling", "Cheese", "Techno"}
	p := Person{
		Name:    "George",
		Age:     40,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)
}
