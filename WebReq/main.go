package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("LCO web Requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type is %T\n", response)

	defer response.Body.Close() //close the connection

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Body is ", string(body))

	// HTTP status code
	fmt.Println("Status Code:", response.StatusCode)

	// Full status
	fmt.Println("Status:", response.Status)
	
}