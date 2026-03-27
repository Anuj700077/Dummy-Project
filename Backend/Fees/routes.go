package fees

import "github.com/gin-gonic/gin"

func FeesRoutes(r *gin.Engine) {
	r.POST("/fees", CreateFees)
	r.GET("/fees", GetLatestFeesHandler)
	r.GET("/fees/student/:sid", GetFeesByStudent)
}
