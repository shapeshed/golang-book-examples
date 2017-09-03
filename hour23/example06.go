package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s := "2006-01-02T15:04:05+07:00"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The hour is %v\n", t.Hour())
	fmt.Printf("The minute is %v\n", t.Minute())
	fmt.Printf("The second is %v\n", t.Second())
	fmt.Printf("The day is %v\n", t.Day())
	fmt.Printf("The month is %v\n", t.Month())
	fmt.Printf("UNIX time is %v\n", t.Unix())
	fmt.Printf("The day of the week is %v\n", t.Weekday())
}
