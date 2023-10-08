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

func CourseIndex(c *gin.Context) {
	employee := c.Query("employee")
	candidate := c.Query("candidate")

	empid, _ := strconv.Atoi(employee)
	canid, _ := strconv.Atoi(candidate)

	var courses []models.Course

	base := db.DB

	if empid > 0 {
		base = base.Where("employee_id = ?", empid)
	} else {
		base = base.Where("candidate_id = ?", canid)
	}

	base.Find(&courses)

	if base.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": base.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courses,
	})
}

func CourseCreate(c *gin.Context) {
	var body models.Course

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	course := models.Course{
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
		Institute:   body.Institute,
		Type:        body.Type,
		DateConduct: body.DateConduct,
		Competency:  body.Competency,
		Venue:       body.Venue,
	}
	result := db.DB.Create(&course)

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

func CourseShow(c *gin.Context) {
	id := c.Param("id")

	var course models.Course
	err := db.DB.First(&course, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": course,
	})
}

func CourseEdit(c *gin.Context) {
	id := c.Param("id")

	var body models.Course

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var course models.Course
	err := db.DB.First(&course, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&course).Updates(models.Course{
		Institute:   body.Institute,
		Type:        body.Type,
		DateConduct: body.DateConduct,
		Competency:  body.Competency,
		Venue:       body.Venue,
		EmployeeID:  body.EmployeeID,
		CandidateID: body.CandidateID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func CourseDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var course models.Course
	err := db.DB.Delete(&course, "ID = ?", id).Error

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
