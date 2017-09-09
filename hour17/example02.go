package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("s", "Hello world", "String help text")
	flag.Parse()
	fmt.Println("value of s:", *s)
}
