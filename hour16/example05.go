package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Guess the name of my pet to win a prize: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if text == "John" {
		fmt.Println("You won! You win chocolate!")
	} else {
		fmt.Println("You didn't win. Better luck next time")
	}
}
