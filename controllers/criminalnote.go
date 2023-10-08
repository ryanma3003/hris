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

func CriminalNoteIndex(c *gin.Context) {
	employee := c.Query("employee")
	candidate := c.Query("candidate")

	empid, _ := strconv.Atoi(employee)
	canid, _ := strconv.Atoi(candidate)

	var criminals []models.CriminalNote

	base := db.DB

	if empid > 0 {
		base = base.Where("employee_id = ?", empid)
	} else {
		base = base.Where("candidate_id = ?", canid)
	}

	base.Find(&criminals)

	if base.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": base.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": criminals,
	})
}

func CriminalNoteCreate(c *gin.Context) {
	var body models.CriminalNote

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	criminal := models.CriminalNote{
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
		Case:        body.Case,
		Description: body.Description,
	}
	result := db.DB.Create(&criminal)

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

func CriminalNoteShow(c *gin.Context) {
	id := c.Param("id")

	var criminal models.CriminalNote
	err := db.DB.First(&criminal, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": criminal,
	})
}

func CriminalNoteEdit(c *gin.Context) {
	id := c.Param("id")

	var body models.CriminalNote

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var criminal models.CriminalNote
	err := db.DB.First(&criminal, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&criminal).Updates(models.CriminalNote{
		Case:        body.Case,
		Description: body.Description,
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func CriminalNoteDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var criminal models.CriminalNote
	err := db.DB.Delete(&criminal, "ID = ?", id).Error

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
