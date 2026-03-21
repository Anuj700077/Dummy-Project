package main

import (
	"github.com/Anuj700077/Dummy-project/database"
	"github.com/Anuj700077/Dummy-project/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Allow React to call backend
	r.Use(cors.Default())

	database.ConnectDB()
	database.CreateTable()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
