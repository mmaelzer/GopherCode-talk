package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var urls []string

func init() {
	urls = []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.yahoo.com",
		"https://www.wikipedia.org",
		"https://golang.org",
	}
}

func main() {
	start := time.Now()
	// START OMIT
	c := make(chan string, len(urls)) // HL

	data := make([]string, len(urls))

	for _, url := range urls {
		go func(url string) {
			response, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			c <- string(body) // HL
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		data[i] = <-c // HL
	}
	// END OMIT
	fmt.Printf("crawled %d urls in %s", len(data), time.Now().Sub(start))
}

// END OMIT
