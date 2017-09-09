package main

import (
	"fmt"
)

func showMemoryAddress(x *int) {
	fmt.Println(x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(&i)
}
