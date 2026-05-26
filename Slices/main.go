package main

import ("fmt"
"sort")
func main() {
	fmt.Println("Slices in Golang")
	var myslice = []int{1,2,34}
	fmt.Printf("%T", myslice)
	fmt.Println("LEngth is ",len(myslice))
	myslice = append(myslice,26)
	fmt.Println(myslice)

	myslice = append(myslice[1:] )
	fmt.Println(myslice)

	highScores := make([]int , 4)
	highScores[0] = 12343
	highScores[1] = 12346
	highScores[2] = 123465
	highScores[3] = 12342
	highScores = append(highScores, 214,512,2345)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	//removing value from slice based on index
	var courses = []string{"Django", "Golang", "Python", "Cprogramming"}
	fmt.Println(courses)
	index  :=  2
	courses = append(courses[:index],courses[index+1:]... )
	fmt.Println(courses)
}
