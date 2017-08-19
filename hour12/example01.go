package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
	c <- "I am not shown until the slowFunc() completes"
}

func main() {
	c := make(chan string)
	go slowFunc(c)
	msg := <-c // Block until we receive a message on the channel
	fmt.Println(msg)
}
