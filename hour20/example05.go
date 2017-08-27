package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	jsonStringData := `{"name":"George","age":40,"hobbies":["Cycling","Cheese","Techno"]}`
	jsonByteData := []byte(jsonStringData)
	p := Person{}
	err := json.Unmarshal(jsonByteData, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}
