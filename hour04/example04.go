package main

import "fmt"

func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}

func main() {
	fmt.Println(sayHi())
}
