package main
import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	p := Person {
		Name : "Asim",
		Age:20,
	}
	fmt.Println(p)
}