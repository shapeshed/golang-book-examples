package main

import (
	"fmt"
)

func echo(s string) {
	fmt.Println(s)
	return
}

func main() {
	s := "Hello World"
	t := "Goodbye Cruel World"
	echo(s)
	echo(t)
}
