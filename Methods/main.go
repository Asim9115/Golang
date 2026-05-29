package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	User1 := User{"Asim", "23co26@gmail.com", true, 21}
	fmt.Println(User1)
	fmt.Printf("Details are %+v\n", User1 )
	fmt.Printf("Name is %v and Email is %v \n", User1.Name, User1.Email )

	User1.GetStatus()
	User1.NewMail()
	fmt.Printf("Name is %v and Email is %v \n", User1.Name, User1.Email )
	User1.PointerNewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user Active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of User is ", u.Email)
}

func (u *User) PointerNewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of User is ", u.Email)
}