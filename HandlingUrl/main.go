package main
import ("fmt"
		"net/url")

const URL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gjf2j4mki"

func main() {
	fmt.Println("Welcome to Url handling in GoLang")
	fmt.Println(URL)
	result, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of query params is %T \n",queryParams) 
	fmt.Println(queryParams["coursename"])
	for _, val := range queryParams {
		fmt.Println("Param is : \n",val)
	}

	partsOfUrl := &url.URL{
		Scheme:"https",
		Host : "asimkhan.me",
		Path: "/sem6",
		RawQuery: "",

	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}