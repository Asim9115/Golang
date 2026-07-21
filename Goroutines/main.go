package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://asimkhan.me", "https://google.com", "https://github.com", "http://localhost",
	}

	for _, url := range urls {
		wg.Add(1)       
		go req(url, &wg)     
	}

	wg.Wait()
}

func req(url string, wg *sync.WaitGroup) {
	defer wg.Done() 

	client := &http.Client{Timeout: 3 * time.Second}

	result, err := client.Get(url)
	if err != nil {
		fmt.Printf("URL: %s \t Error: %v\n", url, err)
		return
	}
	defer result.Body.Close() 

	fmt.Printf("URL: %s \t Status: %s\n", url, result.Status)
}
