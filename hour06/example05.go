package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	cheeses[2] = "Camembert"
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(cheeses)
}
