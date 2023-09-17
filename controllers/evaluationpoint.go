package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func EvaluationPointIndex(c *gin.Context) {
	var evalpoints []models.EvaluationPoint
	err := db.DB.Find(&evalpoints).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": evalpoints,
	})
}

func EvaluationPointCreate(c *gin.Context) {
	var body models.EvaluationPoint

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

func EvaluationPointShow(c *gin.Context) {
	id := c.Param("id")

	var evalpoint models.EvaluationPoint
	err := db.DB.First(&evalpoint, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": evalpoint,
	})
}

func EvaluationPointUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.EvaluationPoint

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var evalpoint models.EvaluationPoint
	err := db.DB.First(&evalpoint, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&evalpoint).Updates(&body)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": evalpoint,
	})
}

func EvaluationPointDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var evalpoint models.EvaluationPoint
	err := db.DB.Delete(&evalpoint, "ID = ?", id).Error

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
