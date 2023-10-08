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

func EducationIndex(c *gin.Context) {
	employee := c.Query("employee")
	candidate := c.Query("candidate")

	empid, _ := strconv.Atoi(employee)
	canid, _ := strconv.Atoi(candidate)

	var educations []models.Education

	base := db.DB

	if empid > 0 {
		base = base.Where("employee_id = ?", empid)
	} else {
		base = base.Where("candidate_id = ?", canid)
	}

	base.Find(&educations)

	if base.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": base.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": educations,
	})
}

func EducationCreate(c *gin.Context) {
	var body models.Education

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	education := models.Education{
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
		Degree:      body.Degree,
		YearComp:    body.YearComp,
		Institute:   body.Institute,
		GradePass:   body.GradePass,
		Subject:     body.Subject,
	}
	result := db.DB.Create(&education)

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

func EducationShow(c *gin.Context) {
	id := c.Param("id")

	var education models.Education
	err := db.DB.First(&education, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": education,
	})
}

func EducationEdit(c *gin.Context) {
	id := c.Param("id")

	var body models.Education

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var education models.Education
	err := db.DB.First(&education, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&education).Updates(models.Education{
		Degree:      body.Degree,
		YearComp:    body.YearComp,
		Institute:   body.Institute,
		GradePass:   body.GradePass,
		Subject:     body.Subject,
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func EducationDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var education models.Education
	err := db.DB.Delete(&education, "ID = ?", id).Error

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
