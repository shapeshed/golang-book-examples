package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Blog string `json:"blog"`
}

func main() {
	var u User
	res, err := http.Get("https://api.github.com/users/shapeshed")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
}
