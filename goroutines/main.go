package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	sites := []string{"https://google.com", "https://facebook.com", "https://golang.org/"}

	for _, site := range sites {
		wg.Add(1)
		go func(s string) {
			res, err := http.Get(s)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(s, ":", res.Status)
			}
			wg.Done()
		}(site)
	}

	wg.Wait()
}
