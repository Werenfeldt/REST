package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Id int `json:"id"`

	Name string `json:"name"`

	Ects float32 `json:"ects"`

	Rating float32 `json:"rating"`
}

var courses = []Course{
	{Id: 2222, Name: "BDSA", Ects: 15, Rating: 34},
	{Id: 4444, Name: "DisSys", Ects: 7.5, Rating: 0},
	{Id: 6666, Name: "IDD", Ects: 7.5, Rating: 12},
	{Id: 7777, Name: "SysDab", Ects: 7.5, Rating: 5},
}

func getCourses(c *gin.Context) {
	c.JSON(http.StatusOK, courses)
}

func newCourse(c *gin.Context) {
	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	courses = append(courses, newCourse)
	c.JSON(http.StatusCreated, newCourse)
}

func getCourse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	for _, v := range courses {
		if v.Id == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}

func deleteCourse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	for i, v := range courses {
		if v.Id == id {
			courses = append(courses[:i], courses[i+1:]...)
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}

func updateCourse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	for _, v := range courses {
		if v.Id == id {
			v = newCourse
			courses = append(courses, v)
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}
