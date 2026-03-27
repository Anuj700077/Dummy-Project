package main

import (
	"github.com/Anuj700077/Dummy-project/database"
	"github.com/Anuj700077/Dummy-project/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	// DB setup
	database.ConnectDB()
	database.CreateTable()
	database.CreateFacultyTable()
	database.CreateMarksTable()
	database.CreateFeeTable()

	// Routes
	routes.SetupRoutes(r)

	// Run server
	r.Run(":8080")
}
