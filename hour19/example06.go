package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	debug := os.Getenv("DEBUG")

	client := &http.Client{
		Timeout: 50 * time.Millisecond,
	}
	request, err := http.NewRequest("GET", "https://ifconfig.co", nil)
	request.Header.Add("Accept", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	if debug == "1" {
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugRequest)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if debug == "1" {
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugResponse)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", body)

}
