package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	go pinger(messages)
	msg := <-messages
	fmt.Println(msg)
}
