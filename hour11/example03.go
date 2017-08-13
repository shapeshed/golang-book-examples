package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("slowFunc() finished")
}

func main() {
	go slowFunc()
	fmt.Println("I am now shown straightaway!")
	time.Sleep(time.Second * 3)
}
