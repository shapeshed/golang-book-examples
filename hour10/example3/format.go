package main

import (
	"fmt"
)

func main() {
	name, role := "Richard Jupp", "Drummer"
	err := fmt.Errorf("The %v %v quit", role, name)
	if err != nil {
		fmt.Println(err)
	}
}
