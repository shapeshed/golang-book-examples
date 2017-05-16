package main

import "fmt"

func getPrize() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}

func main() {
	quantity, prize := getPrize()
	fmt.Printf("You won %v %v\n", quantity, prize)
}
