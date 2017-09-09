package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	for i, n := range numbers {
		fmt.Println("The index of the loop is", i)
		fmt.Println("The value from the slice is", n)
	}
}
