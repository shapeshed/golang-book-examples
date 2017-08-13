package main

import (
	"fmt"
)

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

func main() {
	t := Triangle{base: 3, height: 1}
	t.changeBase(4)
	fmt.Println(t.base)
}
