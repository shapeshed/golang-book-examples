package main

import "fmt"

func main() {
	defer fmt.Println("I am run after the function returns")
	fmt.Println("Hello World!")
}
