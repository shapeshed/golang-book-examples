package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"github.com/shapeshed/darksky"
	"log"
)

func main() {
	s := "Ti esrever dna ti pilf nwod gniht ym tup I"
	fmt.Println(stringutil.Reverse(s))

	params := darksky.RequestParams{
		Key:       "17b1e8cae7b654290659b438557def7e",
		Latitude:  52.847875,
		Longitude: -0.664397,
		Units:     "si",
	}

	forecast, err := darksky.Get(&params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(forecast.Currently.Temperature)
}
