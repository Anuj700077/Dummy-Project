package marks

import "github.com/gin-gonic/gin"

func MarksRoutes(r *gin.Engine) {
	r.POST("/marks", CreateMarks)
	r.GET("/marks", GetMarks)
	r.PUT("/marks", UpdateMarks)
	r.DELETE("/marks/:id", DeleteMarksHandler)
}
