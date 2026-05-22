package main
import "fmt"

type Person struct{
	Name string 
	Age int
}

func getData() (int, string) {
	return 100, "Hello"
}

func main() {
	arr := [3]int{1,2,3}
	fmt.Println("array",arr)

	slice := []int{1,2,3}
	slice = append(slice,4)
	fmt.Println("slice",slice)

	dict := map[string]int{
		"apple":10,
		"banana":20,
	}
	fmt.Println(dict)
	dict["mango"] = 20
	fmt.Println(dict)
	delete(dict,"apple")
	fmt.Println(dict)

	// set := make(map[string]struct{})
	var temp string
	fmt.Println(temp)
}