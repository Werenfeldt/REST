package Server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

var students = []Student{
	{Id: 123, FirstName: "Sofie", LastName: "Go", Email: "Sofie.go@gomail.com", Phone: "123456"},
	{Id: 1666, FirstName: "Sofie", LastName: "Go", Email: "Sofie.go@gomail.com", Phone: "123456"},
}

func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func newStudents(c *gin.Context) {
	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func deleteStudents(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	for i, v := range students {
		if v.Id == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, v)
			// students[i] = students[len(students)-1]
			// students[len(students)-1] = nil
			// students = students[:len(students)-1]
		} else {
			c.JSON(http.StatusNotFound, gin.H{"message": "Student not found"})
		}
	}
}

func updateStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}
