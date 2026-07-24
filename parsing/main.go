package main

import (
	"fmt"
	"regexp"
)

func main() {
	urls := []string{
		"https://google.com", "http://localhost:8080", "https://asimkhan.me", "git@github", "192.168.1.1",
	}
	fmt.Printf("parsing urls : %v\n", urls)
	
	ParseUrl(urls)

}

func ParseUrl(urls []string) {
	// Compile the regex pattern to match any of the characters: :, /, ., or @
	// The '+' sign collapses consecutive matching characters together
	result := make(map[string][]string)
	reg := regexp.MustCompile(`[:/@.]+`)

	for _, url := range urls {
		res := reg.Split(url, -1)
		for _, i := range res {
			result[url] = append(result[url], i)
		}
		
	}
	for key, val := range result{
		fmt.Printf("key: %v, values: %v\n", key, val)
	}

}
