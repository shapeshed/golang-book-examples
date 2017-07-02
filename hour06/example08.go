package main

import "fmt"

func main() {
	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26
	delete(players, "cook")
	fmt.Println(players)
}
