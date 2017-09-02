package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Name   string
	Awake  bool
	Hungry bool
}

func main() {
	c := Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}
