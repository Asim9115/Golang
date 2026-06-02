package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type course struct{
	Name string `json:"coursename"`
	Price int
}

func rmain(){
	r := mux.NewRouter()
	r.HandleFunc("/",GetCourses).Methods("GET")
	fmt.Println("Server Running on http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", r))


}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	mycourse := []course{		
		{"Reactjs",299},
		{"Angularjs",199},
		{"FullStack",499},
}

	finalJson, err := json.Marshal(mycourse)
	if err != nil {
		panic(err)
	}
    w.Header().Set("Content-Type", "application/json")
    w.Write(finalJson)
}