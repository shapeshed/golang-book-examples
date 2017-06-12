package main

import "fmt"

func addition(x int, y int) int {
	return x + y
}

func main() {
	var s string = "three"
	fmt.Println(addition(1, s))
}
