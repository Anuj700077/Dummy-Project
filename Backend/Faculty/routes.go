package faculty

import "github.com/gin-gonic/gin"

func FacultyRoutes(r *gin.Engine) {
	r.POST("/faculty", CreateFaculty)
	r.GET("/faculty", GetFaculty)
	r.PUT("/faculty/:id", UpdateFaculty)
	r.DELETE("/faculty/:id", DeleteFacultyHandler)
}
