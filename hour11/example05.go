package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func responseTime(url string) {

	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	elapsed := time.Since(start).Seconds()

	fmt.Printf("%s took %v seconds \n", url, elapsed)

}

func main() {
	urls := make([]string, 3)
	urls[0] = "https://www.usa.gov/"
	urls[1] = "https://www.gov.uk/"
	urls[2] = "http://www.gouvernement.fr/"

	for _, u := range urls {
		go responseTime(u)
	}

	time.Sleep(time.Second * 5)

}
