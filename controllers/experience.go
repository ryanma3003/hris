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

func ExperienceIndex(c *gin.Context) {
	employee := c.Query("employee")
	candidate := c.Query("candidate")

	empid, _ := strconv.Atoi(employee)
	canid, _ := strconv.Atoi(candidate)

	var experiences []models.Experience

	base := db.DB

	if empid > 0 {
		base = base.Where("employee_id = ?", empid)
	} else {
		base = base.Where("candidate_id = ?", canid)
	}

	base.Find(&experiences)

	if base.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": base.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": experiences,
	})
}

func ExperienceCreate(c *gin.Context) {
	var body models.Experience

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	experience := models.Experience{
		EmployeeID:    body.EmployeeID,
		CandidateID:   body.CandidateID,
		CompanyName:   body.CompanyName,
		CompJoinDate:  body.CompJoinDate,
		CompLeaveDate: body.CompLeaveDate,
		PositionTitle: body.PositionTitle,
		StatusEmp:     body.StatusEmp,
		Thp:           body.Thp,
		Gapok:         body.Gapok,
		Allowance:     body.Allowance,
		OtherBenefit:  body.OtherBenefit,
	}
	result := db.DB.Create(&experience)

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

func ExperienceShow(c *gin.Context) {
	id := c.Param("id")

	var experience models.Experience
	err := db.DB.First(&experience, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": experience,
	})
}

func ExperienceEdit(c *gin.Context) {
	id := c.Param("id")

	var body models.Experience

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var experience models.Experience
	err := db.DB.First(&experience, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&experience).Updates(models.Experience{
		CompanyName:   body.CompanyName,
		CompJoinDate:  body.CompJoinDate,
		CompLeaveDate: body.CompLeaveDate,
		PositionTitle: body.PositionTitle,
		StatusEmp:     body.StatusEmp,
		Thp:           body.Thp,
		Gapok:         body.Gapok,
		Allowance:     body.Allowance,
		OtherBenefit:  body.OtherBenefit,
		EmployeeID:    body.EmployeeID,
		CandidateID:   body.CandidateID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func ExperienceDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var experience models.Experience
	err := db.DB.Delete(&experience, "ID = ?", id).Error

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
