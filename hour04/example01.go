package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}

func main() {
	fmt.Printf("%v\n", isEven(1))
	fmt.Printf("%v\n", isEven(2))
}
