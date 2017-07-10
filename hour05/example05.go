package main

import "fmt"

func main() {
	b := true
	if b {
		fmt.Println("I am executed!")
	} else if b {
		fmt.Println("I am not executed!")
	}
}
