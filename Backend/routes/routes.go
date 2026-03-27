package routes

import (
	"github.com/Anuj700077/Dummy-project/faculty"
	"github.com/Anuj700077/Dummy-project/fees"
	"github.com/Anuj700077/Dummy-project/marks"
	"github.com/Anuj700077/Dummy-project/students"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	students.StudentRoutes(r)
	faculty.FacultyRoutes(r)
	marks.MarksRoutes(r)
	fees.FeesRoutes(r)
}
