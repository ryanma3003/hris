package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func LoanIndex(c *gin.Context) {
	var dep []models.Loan
	err := db.DB.Find(&dep).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dep,
	})
}

func LoanCreate(c *gin.Context) {
	var body models.Loan

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	ln := models.Loan{
		LoanID:     body.LoanID,
		EmployeeID: body.EmployeeID,
		Reason:     body.Reason,
		Amount:     body.Amount}
	result := db.DB.Create(&ln)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func LoanShow(c *gin.Context) {
	id := c.Param("id")

	var ln models.Loan
	err := db.DB.First(&ln, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ln,
	})
}

func LoanUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Loan

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var ln models.Loan
	err := db.DB.First(&ln, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&ln).Updates(models.Loan{
		LoanID:     body.LoanID,
		EmployeeID: body.EmployeeID,
		Reason:     body.Reason,
		Amount:     body.Amount,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": ln,
	})
}

func LoanDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var ln models.Loan
	err := db.DB.Delete(&ln, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
