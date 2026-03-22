package routes

import (
	"github.com/Anuj700077/Dummy-project/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/students", handlers.CreateStudent)     //Creating a students
	r.GET("/students", handlers.GetStudents)        //fetching students
	r.PUT("/students/:id", handlers.UpdateStudents) //updating students by id
	r.DELETE("/students/:id", handlers.DeleteStudent) //deleting students by id
}
