package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to making Web Requests")
	PerfromGetRequest()
}

func PerfromGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code is : ", response.StatusCode)
	fmt.Println("Conetent Length is : ",response.ContentLength)
	// content, err :=ioutil.ReadAll(response.Body) //Byte Format
	// fmt.Println(response)
	// fmt.Println("Content is : ",string(content))

	var responseString strings.Builder
	content, err :=ioutil.ReadAll(response.Body) //Byte Format
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

}