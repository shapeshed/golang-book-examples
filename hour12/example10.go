package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func responseTime(url string, c chan string) {

	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	elapsed := time.Since(start).Seconds()

	c <- fmt.Sprintf("%s took %v seconds \n", url, elapsed)

}

func main() {
	c := make(chan string)

	urls := make([]string, 3)
	urls[0] = "https://www.usa.gov/"
	urls[1] = "https://www.gov.uk/"
	urls[2] = "http://www.gouvernement.fr/"

	for _, u := range urls {
		go responseTime(u, c)
	}

	timeout := time.After(300 * time.Millisecond)

	for i := 1; i <= len(urls); i++ {
		select {
		case res := <-c:
			fmt.Println(res)
		case <-timeout:
			fmt.Println("timeout reached")
			return
		}
	}

}
