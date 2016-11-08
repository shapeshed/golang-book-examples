package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Something went wrong")
	if err != nil {
		fmt.Println(err)
	}
}
