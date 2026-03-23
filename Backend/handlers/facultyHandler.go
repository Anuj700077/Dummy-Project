package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anuj700077/Dummy-project/models"
	"github.com/gin-gonic/gin"
)

func CreateFaculty(c *gin.Context) {
	var faculty models.Faculty
	if err := c.BindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := faculty.CreateFaculty()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Faculty added successfully"})
}

func GetFaculty(c *gin.Context) {
	faculty, err := models.GetAllFaculty()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}

	c.JSON(http.StatusOK, faculty)
}


func UpdateFaculty(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var faculty models.Faculty
	if err := c.BindJSON(&faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = faculty.UpdateFaculty(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Faculty updated successfully"})
}



func DeleteFaculty(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	err = models.DeleteFaculty(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Faculty deleted successfully"})
}
