package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func JobDescriptionIndex(c *gin.Context) {
	var jobdesc []models.JobDescription
	err := db.DB.Find(&jobdesc).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jobdesc,
	})
}

func JobDescriptionCreate(c *gin.Context) {
	var body models.JobDescription

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	jobdesc := models.JobDescription{Name: body.Name, Description: body.Description}
	result := db.DB.Create(&jobdesc)

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

func JobDescriptionShow(c *gin.Context) {
	id := c.Param("id")

	var jobdesc models.JobDescription
	err := db.DB.First(&jobdesc, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jobdesc,
	})
}

func JobDescriptionUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.JobDescription

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var jobdesc models.JobDescription
	err := db.DB.First(&jobdesc, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&jobdesc).Updates(models.JobDescription{
		Name:        body.Name,
		Description: body.Description,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": jobdesc,
	})
}

func JobDescriptionDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var jobdesc models.JobDescription
	err := db.DB.Delete(&jobdesc, "ID = ?", id).Error

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
