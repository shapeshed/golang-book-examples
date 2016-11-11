package main

import "fmt"

func main() {
	const c string = "I ain't gonna gettin' on no airplane, Hannibal!"
	fmt.Println(c)
	var s string
	s = "We are not going on a plane"
	fmt.Println(s)
	s = "We are going on a plane"
	fmt.Println(s)
	// The value of a constant cannot be reassigned.
	// This will cause a compile time error
	// c = "Shut up fool!"
	// fmt.Println(c)
}
