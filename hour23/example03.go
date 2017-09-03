package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("You have two seconds to calculate 19 * 4")
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Time's up! The answer is 74. Did you get it?")
			return
		}
	}

}
