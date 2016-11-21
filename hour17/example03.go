package main

import (
	"flag"
	"fmt"
)

func main() {
	s := flag.String("s", "Hello world", "String help text")
	i := flag.Int("i", 1, "Int help text")
	b := flag.Bool("b", false, "Bool help text")
	flag.Parse()
	fmt.Println("value of s:", *s)
	fmt.Println("value of i:", *i)
	fmt.Println("value of b:", *b)
}
