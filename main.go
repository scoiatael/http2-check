package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.Client{
		// InsecureTLSDial is temporary and will likely be
		// replaced by a different API later.
		Transport: &http2.Transport{},
	}

	resp, err := client.Get("https://http2.golang.org")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
