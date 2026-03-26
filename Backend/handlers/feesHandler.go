package handlers

import (
	"fmt"
	"net/http"

	"github.com/Anuj700077/Dummy-project/models"
	"github.com/gin-gonic/gin"
)


func CreateFees(c *gin.Context) {

	var fee models.Fees

	if err := c.ShouldBindJSON(&fee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err := models.CreateFee(fee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "✅ fee added successfully"})
}

func GetLatestFees(c *gin.Context) {

	fees, err := models.GetLatestFees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch latest fees",
		})
		return
	}

	c.JSON(http.StatusOK, fees)
}


func GetFeesByStudent(c *gin.Context) {

	sid := c.Param("sid")

	var studentID int64
	fmt.Sscan(sid, &studentID)

	fees, err := models.GetFeesByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch student history",
		})
		return
	}

	c.JSON(http.StatusOK, fees)
}
