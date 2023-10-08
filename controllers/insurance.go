package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func InsuranceIndex(c *gin.Context) {
	var insurance []models.Insurance
	err := db.DB.Find(&insurance).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": insurance,
	})
}

func InsuranceCreate(c *gin.Context) {
	var body models.Insurance

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Inisialisasi nilai-nilai default
	defaultInteger := 0

	if body.Verified1 == nil {
		body.Verified1 = &defaultInteger
	}
	if body.Verified2 == nil {
		body.Verified2 = &defaultInteger
	}
	if body.Status == nil {
		body.Status = &defaultInteger
	}

	ins := models.Insurance{
		EmployeeID: body.EmployeeID,
		Reason:     body.Reason,
		Verified1:  body.Verified1,
		Verified2:  body.Verified2,
		Status:     body.Status}
	result := db.DB.Create(&ins)

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

func InsuranceShow(c *gin.Context) {
	id := c.Param("id")

	var ins models.Insurance
	err := db.DB.First(&ins, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ins,
	})
}

func InsuranceUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Insurance

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var ins models.Insurance
	err := db.DB.First(&ins, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&ins).Updates(models.Insurance{
		Verified1: body.Verified1,
		Verified2: body.Verified2,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": ins,
	})
}

func InsuranceDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var ins models.Insurance
	err := db.DB.Delete(&ins, "ID = ?", id).Error

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
