package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake db
var courses []Course

//middleware, Helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	
}

//Controllers - file

//Serve home Route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Api By LCO </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "applcation/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request) {
	 fmt.Println("Get One course")
	w.Header().Set("Content-Type", "applcation/json")
	 
	//Grab Id from Request
	params := mux.Vars(r) 
	//Looping Courses
	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("NO Course Found with Given id %s", params["id"]),)
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Course")
	w.Header().Set("Content-Type", "applcation/json")
	 //Empty Body
	 if r.Body == nil {
		json.NewEncoder(w).Encode("Send the Data")
	 }
	 //{}
	 var course Course
	 _ = json.NewDecoder(r.Body).Decode(&course)
	 if course.IsEmpty() {
		json.NewEncoder(w).Encode("Send all data")
		return
	 }
	 
	 //generating unique id, String
	 //append to courses

	 rand.Seed(time.Now().UnixNano())
	 course.CourseId = strconv.Itoa(rand.Intn(100))
	 courses = append(courses, course)
	 json.NewEncoder(w).Encode(course)
	 return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating Course")
	w.Header().Set("Content-Type", "applcation/json")
	//grab id from request

	params := mux.Vars(r)

	//loop, id remove,add with myid
	for index, course :=  range courses {
		if params["id"] == course.CourseId {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting A Course")
	w.Header().Set("Content-Type", "applcation/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]... )
			json.NewEncoder(w).Encode("Course Deleted Successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return
}

