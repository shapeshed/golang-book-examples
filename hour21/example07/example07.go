package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Name   string `json:"name"`
	Awake  bool   `json:"awake"`
	Hungry bool   `json:"hungry"`
}

func main() {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	c := Config{}
	err = json.Unmarshal(f, &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}
