package main

import (
	"fmt"
)

// Half takes an integer and returns half the value
func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

func main() {
	n, err := Half(19)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}
