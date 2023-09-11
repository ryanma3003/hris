package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func EmployeeIndex(c *gin.Context) {
	var employees []models.Employee
	err := db.DB.Find(&employees).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employees,
	})
}

func EmployeeCreate(c *gin.Context) {
	// Get data req
	var body struct {
		Name  string
		Email string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Create
	employee := models.Employee{Name: body.Name, Email: body.Email}
	result := db.DB.Create(&employee)

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

func EmployeeShow(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	err := db.DB.First(&employee, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

func EmployeeUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body struct {
		Name  string
		Email string
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the data
	var employee models.Employee
	err := db.DB.First(&employee, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&employee).Updates(models.Employee{
		Name:  body.Name,
		Email: body.Email,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

func EmployeeDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var employee models.Employee
	err := db.DB.Delete(&employee, "ID = ?", id).Error

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
