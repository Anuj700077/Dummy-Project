package handlers

import (
	"fmt"
	"net/http"

	database "github.com/Anuj700077/Dummy-project/database"
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
	_, err := database.DB.Exec(
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

func GetStudents(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, sname, fname, address, dob FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	defer rows.Close()
	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.Sname, &student.Fname, &student.Address, &student.Dob)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Students fetching error"})
			return
		}
		students = append(students, student)
	}
	c.JSON(http.StatusOK, students)
}

func UpdateStudents(c *gin.Context) {

	// get ID from URL
	id := c.Param("id")

	var student models.Student

	// get updated data from frontend
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	// update query
	_, err := database.DB.Exec(
		`UPDATE students 
		 SET sname=$1, fname=$2, address=$3, dob=$4 
		 WHERE id=$5`,
		student.Sname,
		student.Fname,
		student.Address,
		student.Dob,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update student",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated successfully",
	})
}

func DeleteStudent(c *gin.Context) {

	id := c.Param("id")

	_, err := database.DB.Exec(
		"DELETE FROM students WHERE id=$1",
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete student",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully",
	})
}
