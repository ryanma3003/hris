package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func DepartmentIndex(c *gin.Context) {
	var dep []models.Department
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

func DepartmentCreate(c *gin.Context) {
	var body models.Department

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	dep := models.Department{Name: body.Name, DivisionID: body.DivisionID}
	result := db.DB.Create(&dep)

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

func DepartmentShow(c *gin.Context) {
	id := c.Param("id")

	var dep models.Department
	err := db.DB.First(&dep, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dep,
	})
}

func DepartmentUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Department

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var dep models.Department
	err := db.DB.First(&dep, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&dep).Updates(models.Department{
		Name:       body.Name,
		DivisionID: body.DivisionID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": dep,
	})
}

func DepartmentDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var dep models.Department
	err := db.DB.Delete(&dep, "ID = ?", id).Error

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
