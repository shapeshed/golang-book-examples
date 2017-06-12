package main

import "fmt"

func sayHello(s string) string {
	return "Hello " + s
}

func main() {
	fmt.Println(sayHello("George"))
}
