package main

import "fmt"

func main() {
	defer fmt.Println("I am run after the function completes")
	fmt.Println("Hello World!")
}
