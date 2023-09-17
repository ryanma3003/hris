package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func EvaluationFormIndex(c *gin.Context) {
	var evalforms []models.EvaluationForm
	err := db.DB.Find(&evalforms).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": evalforms,
	})
}

func EvaluationFormCreate(c *gin.Context) {
	var body models.EvaluationForm

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	result := db.DB.Create(&body)

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

func EvaluationFormShow(c *gin.Context) {
	id := c.Param("id")

	var evalform models.EvaluationForm
	err := db.DB.First(&evalform, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": evalform,
	})
}

func EvaluationFormUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.EvaluationForm

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var evalform models.EvaluationForm
	err := db.DB.First(&evalform, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&evalform).Updates(&body)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": evalform,
	})
}

func EvaluationFormDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var evalform models.EvaluationForm
	err := db.DB.Delete(&evalform, "ID = ?", id).Error

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
