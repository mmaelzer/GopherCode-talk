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
		"https://socialcode.com",
		"https://www.wikipedia.org",
		"https://golang.org",
	}
}

func main() {
	start := time.Now()
	// START OMIT
	data := []string{} // HL

	for _, url := range urls {
		response, err := http.Get(url) // HL
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close() // HL

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, string(body)) // HL
	}
	// END OMIT
	// fmt.Println(data)
	fmt.Printf("crawled %d urls in %s", len(data), time.Now().Sub(start))
}

// END OMIT
