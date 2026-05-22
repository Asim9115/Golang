package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	fmt.Println("Taking Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Something ")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("The Input is,", input) 
	fmt.Printf("Type is %T", input) 
}