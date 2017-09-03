package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://petition.parliament.uk/petitions")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	src := string(body)

	re := regexp.MustCompile("\\<h3\\>.*\\</h3\\>")
	titles := re.FindAllString(src, -1)

	for _, title := range titles {
		fmt.Println(title)
	}

}
