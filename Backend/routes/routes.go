package routes

import (
	"github.com/Anuj700077/Dummy-project/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/students", handlers.CreateStudent)
}
