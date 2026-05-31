package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
	
)

func main() {
	fmt.Println("Welcome to making Web Requests")
	//PerfromGetRequest()
	// PerformPostRequest()
	SendFormData()
}

func SendFormData() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("Name","ASim")
	data.Add("Age", "20")
	data.Add("Email","asim@asim.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	defer response.Body.Close()
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
	{
		"cousename" : "Go with GoLang",
		"price" : 12,
		"platform" : "lco.in"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	content , err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	fmt.Println(response.StatusCode)
		defer response.Body.Close()

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