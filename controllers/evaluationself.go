package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func SelfPerformanceIndex(c *gin.Context) {
	var selfperfs []models.SelfPerformance
	err := db.DB.Find(&selfperfs).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": selfperfs,
	})
}

func SelfPerformanceCreate(c *gin.Context) {
	var body models.SelfPerformance

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

func SelfPerformanceShow(c *gin.Context) {
	id := c.Param("id")

	var selfperf models.SelfPerformance
	err := db.DB.First(&selfperf, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": selfperf,
	})
}

func SelfPerformanceUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.SelfPerformance

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var selfperf models.SelfPerformance
	err := db.DB.First(&selfperf, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&selfperf).Updates(&body)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": selfperf,
	})
}

func SelfPerformanceDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var selfperf models.SelfPerformance
	err := db.DB.Delete(&selfperf, "ID = ?", id).Error

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
