package fees

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CREATE
func CreateFees(c *gin.Context) {

	var fee Fees

	if err := c.ShouldBindJSON(&fee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	err := CreateFee(fee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "fee added successfully"})
}

// GET LATEST
func GetLatestFeesHandler(c *gin.Context) {

	feesList, err := GetLatestFees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch latest fees",
		})
		return
	}

	c.JSON(http.StatusOK, feesList)
}

// GET BY STUDENT
func GetFeesByStudent(c *gin.Context) {

	sidParam := c.Param("sid")

	sid, err := strconv.ParseInt(sidParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid student id"})
		return
	}

	feesList, err := GetFeesByStudentID(sid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch student history",
		})
		return
	}

	c.JSON(http.StatusOK, feesList)
}
