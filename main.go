package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



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

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Teacher *Teacher `json:"teacher,omitempty"`

	Ects float32 `json:"ects,omitempty"`

	Rating float32 `json:"rating,omitempty"`
}

type Teacher struct {

	Id int64 `json:"id,omitempty"`

	Course *Course `json:"Course,omitempty"`

	Firstname string `json:"firstname,omitempty"`

	Lastname string `json:"lastname,omitempty"`
}

var students = []Student{
	{Id: 123, FirstName: "Sofie", LastName: "Go", Email: "Sofie.go@gomail.com", Phone: "123456"},
	{Id: 1666, FirstName: "Sofie", LastName: "Go", Email: "Sofie.go@gomail.com", Phone: "123456"},
}

func main(){
	router := gin.Default()
	router.GET("/students", getStudents)
	//router.PUT("/students", updateStudents)
	router.POST("/students", newStudents)
	router.DELETE("/students/:id", deleteStudents)

	router.Run("localhost:8080")
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


