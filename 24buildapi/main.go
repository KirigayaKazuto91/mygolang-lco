package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"mygolang-lco/23mymodule/vendor/github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

// Model for course - file
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"` // Pointer to Author struct
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

// Controller - file
// Serve home route

func serveHome (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to course API</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Get the course id from the URL
	params := mux.Vars(r)

	// Loop through all the courses and find the course with the given id
	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// If no course found with the given id
	json.NewEncoder(w).Encode("Course not found")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if the request body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("Please send valid data")
		return
	}

	// generate a random id for the course
	// append the course to the courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}
