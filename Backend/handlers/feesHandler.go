package handlers

import (
	"net/http"

	"github.com/Anuj700077/Dummy-project/models"
	"github.com/gin-gonic/gin"
)

// CREATE
func CreateFees(c *gin.Context) {

	var fee models.Fees

	if err := c.ShouldBindJSON(&fee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "❌ invalid input"})
		return
	}

	err := models.CreateFee(fee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "✅ fee added successfully"})
}

// GET ALL
func GetAllFees(c *gin.Context) {

	fees, err := models.GetAllFees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ could not fetch fees",
		})
		return
	}

	c.JSON(http.StatusOK, fees)
}
