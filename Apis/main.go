package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Println("API")
	r := mux.NewRouter()
	//seeding

courses = append(courses,
	Course{
		CourseId:    "1",
		CourseName:  "Golang",
		CoursePrice: 999,
		Author: &Author{
			Fullname: "Asim Khan",
			Website:  "asimkhan.me",
		},
	},
	Course{
		CourseId:    "2",
		CourseName:  "Python",
		CoursePrice: 799,
		Author: &Author{
			Fullname: "John Doe",
			Website:  "johndoe.dev",
		},
	},
	Course{
		CourseId:    "3",
		CourseName:  "JavaScript",
		CoursePrice: 899,
		Author: &Author{
			Fullname: "Jane Smith",
			Website:  "janesmith.io",
		},
	},
	Course{
		CourseId:    "4",
		CourseName:  "Java",
		CoursePrice: 1099,
		Author: &Author{
			Fullname: "Rahul Sharma",
			Website:  "rahulcodes.com",
		},
	},
	Course{
		CourseId:    "5",
		CourseName:  "Rust",
		CoursePrice: 1299,
		Author: &Author{
			Fullname: "Alice Brown",
			Website:  "alicebrown.dev",
		},
	},
)

//Routing
r.HandleFunc("/",serveHome).Methods("GET")
r.HandleFunc("/courses", getAllCourses).Methods("GET")
r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
r.HandleFunc("/course",createOneCourse).Methods("POST")
r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))
	
}

//Controllers - file

//Serve home Route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Api By LCO </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter,r *http.Request) {
	 fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")
	 
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
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
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

