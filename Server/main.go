package main

import (
	"github.com/gin-gonic/gin"
)

// curl -i -X POST -H "Content-Type: application/json" -d "{"id": "8888", "name": "DMat"}" http://localhost:8080/courses

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.POST("/students", newStudents)
	router.DELETE("/students/:id", deleteStudents)

	router.GET("courses", getCourses)
	router.GET("courses/:id", getCourse)
	router.POST("courses", newCourse)
	router.DELETE("courses/:id", deleteCourse)
	router.PUT("courses/:id", updateCourse)

	router.Run("localhost:8080")
}
