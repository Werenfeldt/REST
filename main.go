package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// curl -i -X POST -H "Content-Type: application/json" -d "{"id": "8888", "name": "DMat"}" http://localhost:8080/courses

type Student struct {
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	//StudyProgram string `json:"StudyProgram,omitempty"`
	//Course []Course `json:"Course,omitempty"`
	//Ects float32 `json:"ects,omitempty"`
}

type Course struct {

	Id int `json:"id"`

	Name string `json:"name"`

	Teacher *Teacher `json:"teacher"`

	Ects float32 `json:"ects"`

	Rating float32 `json:"rating"`
}

type Teacher struct {

	Id int64 `json:"id,omitempty"`

	//Course *Course `json:"Course,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`
}

var students = []Student{
	{Id: 123, FirstName: "Sofie", LastName: "Go", Email: "Sofie.go@gomail.com", Phone: "123456"},
	{Id: 1666, FirstName: "Sofie", LastName: "Go", Email: "Sofie.go@gomail.com", Phone: "123456"},
}

var teacher1 = Teacher{Id: 11, Firstname: "Hanne", Lastname: "Hansen" } 
var teacher2 = Teacher{Id: 33, Firstname: "Børge", Lastname: "Bentsen" }
var teacher3 = Teacher{Id: 55, Firstname: "Roberto", Lastname: "Robertsen" }

var courses = []Course{
	{Id: 2222, Name: "BDSA", Teacher: &teacher1, Ects: 15, Rating: 34},
	{Id: 4444, Name: "DisSys", Teacher: &teacher2, Ects: 7.5, Rating: 0},
	{Id: 6666, Name: "IDD", Teacher: &teacher3, Ects: 7.5, Rating: 12},
	{Id: 7777, Name: "SysDab", Teacher: &teacher1, Ects: 7.5, Rating: 5}, 
}

func main(){
	router := gin.Default()
	router.GET("/students", getStudents)
	//router.PUT("/students", updateStudents)
	router.POST("/students", newStudents)
	router.DELETE("/students/:id", deleteStudents)

	router.GET("courses", getCourses)
	router.GET("courses/:id", getCourse)
	router.POST("courses/", newCourse)
	router.DELETE("courses/:id", deleteCourse)
	//router.PUT("courses", getCourses)

	router.Run("localhost:8080")
}

func getCourses(c *gin.Context){
	c.IndentedJSON(http.StatusOK, courses)
}

func newCourse(c *gin.Context){
	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourse(c *gin.Context){
	id, _ := strconv.Atoi(c.Params.ByName("id")) 

	for _, v:= range courses{
		if v.Id==id{
			c.IndentedJSON(http.StatusOK, v)			
		} 
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Course not found"})
}

func deleteCourse(c *gin.Context){
	id, _ := strconv.Atoi(c.Params.ByName("id")) 

	for i, v := range courses{
		if v.Id == id {
			courses = append(courses[:i], courses[i+1:]... )
			c.IndentedJSON(http.StatusOK, courses)
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Course not found"})
		}
	}
}

func getStudents(c *gin.Context){
	c.IndentedJSON(http.StatusOK, students)
}

func newStudents(c *gin.Context){
	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func deleteStudents(c *gin.Context){
	id, _ := strconv.Atoi(c.Params.ByName("id")) 

	for i, v := range students{
		if v.Id == id {
			students = append(students[:i], students[i+1:]... )
			c.IndentedJSON(http.StatusOK, v)
			// students[i] = students[len(students)-1]
			// students[len(students)-1] = nil
			// students = students[:len(students)-1]
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Student not found"})
		}
	}
}

func updateStudents(c *gin.Context){
	c.IndentedJSON(http.StatusOK, students)
}


