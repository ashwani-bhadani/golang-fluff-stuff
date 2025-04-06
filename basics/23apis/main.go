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

// model for course -file
type Course struct {
	CourseId    string  `json:"coruseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float32 `json:"price"`
	Author      *Author `json:"author"` //we want to pass the ref, so declare as pointer
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake data served from here
var courses []Course

// helper functions - in separate file, since this method is part of struct,
// need to pass pointer to struct before the method name, return type is bool
func (c *Course) IsEmpty() bool {
	//from pointer I get acess to entire object
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" //as we will self generate a course id
}

// controller - separate file
// serve home route , governed by 2 reader from where you get value & write to where you send response
func serveHome(w http.ResponseWriter, r *http.Request) { //common practice to declare this way
	w.Write([]byte("<h1>Welcome to API tutorials with Golang!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //whatever is inside encode will be returned as a JSON
}

// good practice of writing api below, func should be named get course only
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	//for one course you must have an id & loop thorugh arr & return
	fmt.Println("Get one course by id.")

	//grab id from request
	params := mux.Vars(r) //extracts all vars from the request
	fmt.Printf("params value : %v & is type %T\n", params, params)

	//loop thru courses, find by id & return, not worried about index
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	/*
		for _, course := range courses
			>This line is looping through a slice named courses.
			>The underscore _ is used to ignore the index of the loop.
			>course is the actual item (likely a struct) from the courses slice in each iteration.
		if course.CourseId == params["id"]
			>This checks if the current course's CourseId matches the value of "id" from the params map.
			>params["id"] is assumed to be a string key-value pair (e.g., from a URL query or route parameter).
	*/

	//below encoder once ecode it called flow auto-stops & returns, so return stmt after this is redundant
	json.NewEncoder(w).Encode("No course found with given id!")
	return //redundant return stmt can also remove
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	//here we have to decode a json
	fmt.Println("create on course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Data Empty error: Please send valid data")
	}

	//what if data is {} as a struct, destructure
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) //passing a ref, don't care about return value
	if course.IsEmpty() {                       //just craft the response if non-empty now
		json.NewEncoder(w).Encode("Data Empty error: The JSON is blank/empty")
		return
	}

	//check & not add if a course already exists
	//loop, title matches the course .course name

	//generate a unique id & convert it to string
	//append course into courses
	// rand.Seed(time.Now().UnixNano()) //deprecated as there's always a need to provide seed value

	//applicable for Go 1.20+, This is cleaner, especially in concurrent code
	// or libraries where you don't want to mess with the global random source.
	var randomInt = rand.New(rand.NewSource(time.Now().UnixNano()))

	course.CourseId = strconv.Itoa(randomInt.Intn(100)) //use the random number & convert it to string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course) //returns the response, & auto exits the method
	// return //redundant
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	//1. grab id from req
	params := mux.Vars(r)

	//2. stesp: loop, id, remove, add with my id again
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			//removing a value from a slice based on index idx as in slices tutorial
			courses = append(courses[:idx], courses[idx+1:]...) //combines portions
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) //decoding json based on struct
			//we need to generate a id, again
			course.CourseId = params["id"]
			courses = append(courses, course) //create new course to return
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: send a response when id not found, or empty {}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	//loop, find by id, remove
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...) //delete item at idx & combine
			break                                               //as deleted item, can stop loop
		}
	}

}

//handling the routes in golang

func main() {
	fmt.Println("Building APIs with golang!")
	r := mux.NewRouter()

	//seeding, adding data into courses
	courses = append(courses, Course{CourseId: "23", CourseName: "React Js", CoursePrice: 499.0, Author: &Author{Fullname: "Ashwani", Website: "github.com"}})
	courses = append(courses, Course{CourseId: "33", CourseName: "Golang", CoursePrice: 699.0, Author: &Author{Fullname: "Ashwani", Website: "github.com"}})
	courses = append(courses, Course{CourseId: "43", CourseName: "Java", CoursePrice: 399.0, Author: &Author{Fullname: "Ashwani", Website: "github.com"}})
	courses = append(courses, Course{CourseId: "53", CourseName: "Spring Boot", CoursePrice: 599.0, Author: &Author{Fullname: "Ashwani", Website: "github.com"}})

	//routing : take the router object & call the handle func, pass on the ref of methods
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") //if you change id to courseId, then use params["courseId"]
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", createOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", createOneCourse).Methods("DEL")

	//listen to a port, log comes handy & helps
	log.Fatal(http.ListenAndServe(":4000", r))
}
