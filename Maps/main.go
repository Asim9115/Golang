package main
import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["CP"] = "C Programming"
	languages["TS"] = "TypeScript"
	fmt.Println(languages)
	fmt.Println("Py is ", languages["PY"])
	//deleting
	delete(languages, "JS")
	fmt.Println(languages)

	//loops
	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}

}
