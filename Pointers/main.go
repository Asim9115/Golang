package main

import "fmt"

func main() {
	fmt.Println("Welocme to pointers")
	var ptr *int
	fmt.Println(ptr)
	myNumber := 26
	var mypointer  = &myNumber
	fmt.Println(mypointer)
	fmt.Println(*mypointer)

	*mypointer = *mypointer *2
	fmt.Println("New value ", myNumber)
}
