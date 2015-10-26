package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	worker("https://www.google.com")
}

// START OMIT
func worker(url string) {
	body, err := crawl(url)
	if err != nil {
		fmt.Printf("Error crawling %s: %s", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", body)
}

func crawl(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// END OMIT
