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

func ReferenceIndex(c *gin.Context) {
	employee := c.Query("employee")
	candidate := c.Query("candidate")

	empid, _ := strconv.Atoi(employee)
	canid, _ := strconv.Atoi(candidate)

	var references []models.Reference

	base := db.DB

	if empid > 0 {
		base = base.Where("employee_id = ?", empid)
	} else {
		base = base.Where("candidate_id = ?", canid)
	}

	base.Find(&references)

	if base.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": base.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": references,
	})
}

func ReferenceCreate(c *gin.Context) {
	var body models.Reference

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	reference := models.Reference{
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
		RefName:     body.RefName,
		RefPhone:    body.RefPhone,
		RefRelation: body.RefRelation,
		RefTitle:    body.RefTitle,
	}
	result := db.DB.Create(&reference)

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

func ReferenceShow(c *gin.Context) {
	id := c.Param("id")

	var reference models.Reference
	err := db.DB.First(&reference, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": reference,
	})
}

func ReferenceEdit(c *gin.Context) {
	id := c.Param("id")

	var body models.Reference

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var reference models.Reference
	err := db.DB.First(&reference, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&reference).Updates(models.Reference{
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
		RefName:     body.RefName,
		RefPhone:    body.RefPhone,
		RefRelation: body.RefRelation,
		RefTitle:    body.RefTitle,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func ReferenceDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var reference models.Reference
	err := db.DB.Delete(&reference, "ID = ?", id).Error

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
