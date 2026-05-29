package main

import (
	"fmt"


)

func main() {
	fmt.Println("Functions")
	result := adder(20, 6)
	fmt.Println("Value is ", result)
	total, mesg := proAdder(1,2,3,4,5,6,5)
	fmt.Printf("total is %v and message is %v", total, mesg)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for i := range values {
		total += values[i]
	}
	return total, "heres the total"
}