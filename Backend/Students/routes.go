package students

import "github.com/gin-gonic/gin"

func StudentRoutes(r *gin.Engine) {
	r.POST("/students", CreateStudent)
	r.GET("/students", GetStudents)
	r.PUT("/students/:id", UpdateStudents)
	r.DELETE("/students/:id", DeleteStudent)
}
