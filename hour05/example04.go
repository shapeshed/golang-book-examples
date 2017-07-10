package main

import "fmt"

func main() {
	b := true
	if b {
		fmt.Println("I am executed!")
	}
	if b {
		fmt.Println("I am also executed!")
	}
}
