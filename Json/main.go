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
	DecodeJson()

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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "Reactjs",
                "Price": 299,
                "website": "LearnCodeonline",
                "tags": ["webdev","js"]
        }
	`)

	var lcoCourse course 
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	//cases where add data to key value non-struct

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", key, val,val)
	}
}