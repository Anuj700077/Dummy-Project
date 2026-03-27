package marks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateMarks(c *gin.Context) {
	var m Marks

	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := m.CreateMark()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save marks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Marks saved successfully"})
}

// GET ALL
func GetMarks(c *gin.Context) {

	marksList, err := GetAllMarks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch marks"})
		return
	}

	c.JSON(http.StatusOK, marksList)
}

// UPDATE
func UpdateMarks(c *gin.Context) {
	var m Marks

	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	err := m.UpdateMark()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "marks updated successfully"})
}

// DELETE
func DeleteMarksHandler(c *gin.Context) {

	idParam := c.Param("id")

	sid, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = DeleteMark(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "marks deleted successfully"})
}
