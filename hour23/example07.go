package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s1 := "2017-09-03T18:00:00+00:00"
	s2 := "2017-09-04T18:00:00+00:00"
	today, err := time.Parse(time.RFC3339, s1)
	if err != nil {
		log.Fatal(err)
	}
	tomorrow, err := time.Parse(time.RFC3339, s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today.After(tomorrow))
	fmt.Println(today.Before(tomorrow))
	fmt.Println(today.Equal(tomorrow))
}
