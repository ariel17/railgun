package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url1 := flag.String("url1", "", "URL as base of comparison")
	url2 := flag.String("url2", "", "URL as content to compare")
	flag.Parse()

	fmt.Printf("url1=%s, url2=%s\n", *url1, *url2)

	body1, err := get(*url1)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("Content length #1: %d\n", len(body1))

	body2, err := get(*url1)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("Content length #2: %d\n", len(body2))
}

func get(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
