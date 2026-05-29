package main
import "fmt"

func main() {
	fmt.Println("Defers in Golang")
	defer fmt.Println("First Defer")
	defer fmt.Println("Second Defer")
	defer fmt.Println("Third Defer")
	fmt.Println("Last Line")
	//LIFO DEFER -> 1, 2, 3  -> Print
	Mydefer()
}

func Mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
// 4
// 3
// 2
// 1
// 0
// Third Defer
// Second Defer
// First Defer