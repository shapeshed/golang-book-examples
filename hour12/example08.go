package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	go sender(messages)

	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
