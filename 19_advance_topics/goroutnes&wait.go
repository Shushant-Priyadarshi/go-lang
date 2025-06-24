package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websites := []string{
		"http://google.com",
		"http://fb.com",
		"http://amazon.com",
		"http://linkedin.com",
		"http://github.com",
		"http://shushantpriyadarshi.me",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)

	}

	wg.Wait()

}

func getStatusCode(website string) {
	res, err := http.Get(website)
	if err != nil {
		fmt.Printf("Cant get to: %s\n", website)
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, website)
	}
	defer wg.Done()
}
