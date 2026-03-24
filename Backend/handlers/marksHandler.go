package handlers

import (
	"net/http"
	"strconv"

	"github.com/Anuj700077/Dummy-project/models"
	"github.com/gin-gonic/gin"
)

func CreateMarks(c *gin.Context) {
	var m models.Marks

	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := m.CreateMark()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Marks saved successfully"})
}

func GetMarks(c *gin.Context) {

	marks, err := models.GetAllMarks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch marks"})
		return
	}

	c.JSON(http.StatusOK, marks)
}



func UpdateMarks(c *gin.Context) {
	var m models.Marks

	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	err := m.UpdateMark()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "marks updated successfully",
	})
}




func DeleteMarks(c *gin.Context) {

	idParam := c.Param("id")

	sid, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = models.DeleteMark(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "marks deleted successfully",
	})
}

