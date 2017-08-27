package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Switch struct {
	On bool `json:"on"`
}

func main() {
	jsonStringData := `{"on":"true"}`
	jsonByteData := []byte(jsonStringData)
	s := Switch{}
	err := json.Unmarshal(jsonByteData, &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", s)
}
