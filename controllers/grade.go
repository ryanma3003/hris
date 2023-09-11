package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func GradeIndex(c *gin.Context) {
	var grades []models.Grade
	err := db.DB.Find(&grades).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": grades,
	})
}

func GradeCreate(c *gin.Context) {
	var body models.Grade

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Create
	grade := models.Grade{GradeId: body.GradeId, Min: body.Min, Max: body.Max, Struktur: body.Struktur}
	result := db.DB.Create(&grade)

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

func GradeShow(c *gin.Context) {
	id := c.Param("id")

	var grade models.Grade
	err := db.DB.First(&grade, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": grade,
	})
}

func GradeUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Grade

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the data
	var grade models.Grade
	err := db.DB.First(&grade, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&grade).Updates(models.Grade{
		GradeId:  body.GradeId,
		Min:      body.Min,
		Max:      body.Max,
		Struktur: body.Struktur,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": grade,
	})
}

func GradeDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var grade models.Grade
	err := db.DB.Delete(&grade, "ID = ?", id).Error

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
