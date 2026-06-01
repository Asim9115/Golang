package main

import (
	"encoding/json"
	"fmt"
	
)

type course struct{
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json Handling in GoLang")
	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs",299, "LearnCodeonline", "123456", []string{"webdev", "js"}},
		{"Angularjs",199, "LearnCodeonline", "abcde", []string{"webdev", "js"}},
		{"FullStack",499, "LearnCodeonline", "123abc", nil},

	}
	//package this as json

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses,"", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}