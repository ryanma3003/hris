package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func HealthDiseaseIndex(c *gin.Context) {
	employee := c.Query("employee")
	candidate := c.Query("candidate")

	empid, _ := strconv.Atoi(employee)
	canid, _ := strconv.Atoi(candidate)

	var healths []models.HealthDisease

	base := db.DB

	if empid > 0 {
		base = base.Where("employee_id = ?", empid)
	} else {
		base = base.Where("candidate_id = ?", canid)
	}

	base.Find(&healths)

	if base.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": base.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": healths,
	})
}

func HealthDiseaseCreate(c *gin.Context) {
	var body models.HealthDisease

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	health := models.HealthDisease{
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
		Name:        body.Name,
		Description: body.Description,
	}
	result := db.DB.Create(&health)

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

func HealthDiseaseShow(c *gin.Context) {
	id := c.Param("id")

	var health models.HealthDisease
	err := db.DB.First(&health, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": health,
	})
}

func HealthDiseaseEdit(c *gin.Context) {
	id := c.Param("id")

	var body models.HealthDisease

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var health models.HealthDisease
	err := db.DB.First(&health, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&health).Updates(models.HealthDisease{
		Name:        body.Name,
		Description: body.Description,
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func HealthDiseaseDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var health models.HealthDisease
	err := db.DB.Delete(&health, "ID = ?", id).Error

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
