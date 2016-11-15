package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 1 // int
	intToString := strconv.Itoa(i)
	var s = " egg" // string
	var breakfast = intToString + s
	fmt.Println(breakfast)
}
