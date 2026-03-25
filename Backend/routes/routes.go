package routes

import (
	"github.com/Anuj700077/Dummy-project/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	
	r.POST("/students", handlers.CreateStudent)
	r.GET("/students", handlers.GetStudents)
	r.PUT("/students/:id", handlers.UpdateStudents)
	r.DELETE("/students/:id", handlers.DeleteStudent)

	
	r.POST("/faculty", handlers.CreateFaculty)
	r.GET("/faculty", handlers.GetFaculty)
	r.PUT("/faculty/:id", handlers.UpdateFaculty)
	r.DELETE("/faculty/:id", handlers.DeleteFaculty)

	
	r.POST("/marks", handlers.CreateMarks)
	r.GET("/marks", handlers.GetMarks)
	r.PUT("/marks", handlers.UpdateMarks)
	r.DELETE("/marks/:id", handlers.DeleteMarks)

	
	r.POST("/fees", handlers.CreateFees)
	r.GET("/fees", handlers.GetAllFees) 
}
