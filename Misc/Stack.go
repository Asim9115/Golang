package main
import (
	"errors"
	"fmt"
)
const MAX = 100
type Stack struct{
	data [MAX]int
	top int
}

func main() {
	fmt.Println("Welcome to data Structure Stack")
	fmt.Println("Stack Intialized")
	s := Stack{}
	s.top = -1
	for {
		var choice int
		fmt.Println("\n1. Push")
		fmt.Println("2. Pop")
		fmt.Println("3. Peek")
		fmt.Println("4. Display")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the element to push")
			var element int
			fmt.Scan(&element)
			s.Push(element)
		case 2:
			element, err := s.Pop()
			if err != nil {
				fmt.Println(err)
				
				
			} else{
			fmt.Printf("Popped element is %d", element)
			}
		case 3:
			element , err := s.Peek()
			if err != nil {
				fmt.Println(err)
			} else {
			fmt.Println("Top Element is ", element)
			}
		
		case 5:
			fmt.Println("Goodbye")
			return
		case 4:
			s.Display()
		default:
		fmt.Println("Invalid choice")
		}
	}
}

func (s *Stack) Push(element int) {
	if s.top == MAX-1 {
		fmt.Println("stack overflow")
		return
	} 
	s.top += 1
		s.data[s.top] = element
		
	
}

func (s *Stack) Pop() (int, error) {
	if s.top == -1 {
		return -1, errors.New("stack underflow: add element first")
	}

	
	data := s.data[s.top]
	s.top--
	return data, nil
}

func (s *Stack) Display(){
	if s.top == -1 {
		fmt.Println("Stack empty")
		return
	}
	for i, value := range s.data[:s.top+1] {
		fmt.Printf("data at index %d is %d\n", i, value)
	}

}

func (s *Stack) Peek() (int , error) {
	if s.top == -1 {
		return -1, errors.New("Stack underflow")
	}
	return s.data[s.top], nil
}