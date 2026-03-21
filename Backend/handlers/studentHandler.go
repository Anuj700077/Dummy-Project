package handlers

import (
	"fmt"
	"net/http"

	databse "github.com/Anuj700077/Dummy-project/database"
	"github.com/Anuj700077/Dummy-project/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student

	// Get data from frontend (JSON)
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert into DB
	_, err := databse.DB.Exec(
		"INSERT INTO students(sname, fname, address, dob) VALUES($1,$2,$3,$4)",
		student.Sname,
		student.Fname,
		student.Address,
		student.Dob,
	)

	fmt.Println("Student Name:", student.Sname)
	fmt.Println("Father Name:", student.Fname)
	fmt.Println("Address:", student.Address)
	fmt.Println("DOB:", student.Dob)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student added successfully"})
}
