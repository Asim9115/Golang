package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to handling file in Golang")
	content := "This is a content that should be put in a file using Golang"
	file, err := os.Create("./mylocgofile.txt")
	CheckNilError(err)

	length, err := io.WriteString(file, content)
	CheckNilError(err)
	fmt.Println("Length is ",length)
	ReadFile("./mylocgofile.txt")
	defer file.Close()
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	CheckNilError(err)
	fmt.Println("Text Data is \n", string(databyte))
}

func CheckNilError(err error) {
	if err != nil {
		panic(err)
	}
}