package main

import (
	"fmt"
)

const greeting string = "Hello, world"

func main() {
	greeting = "Goodbye, cruel world"
	fmt.Println(greeting)
}
