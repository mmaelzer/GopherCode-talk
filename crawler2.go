package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup // HL

	data := make([]string, len(urls))

	for i, url := range urls {
		wg.Add(1)                    // HL
		go func(i int, url string) { // HL
			response, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			data[i] = string(body)
			wg.Done() // HL
		}(i, url)
	}
	wg.Wait() // HL
	// END OMIT
	fmt.Printf("crawled %d urls in %s", len(data), time.Now().Sub(start))
}

// END OMIT
