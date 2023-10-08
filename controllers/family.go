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

func FamilyIndex(c *gin.Context) {
	employee := c.Query("employee")
	candidate := c.Query("candidate")

	empid, _ := strconv.Atoi(employee)
	canid, _ := strconv.Atoi(candidate)

	var familys []models.Family

	base := db.DB

	if empid > 0 {
		base = base.Where("employee_id = ?", empid)
	} else {
		base = base.Where("candidate_id = ?", canid)
	}

	base.Find(&familys)

	if base.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": base.Error,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": familys,
	})
}

func FamilyCreate(c *gin.Context) {
	var body models.Family

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	family := models.Family{EmployeeID: body.EmployeeID, CandidateID: body.CandidateID, FamRelation: body.FamRelation, FamName: body.FamName, FamProfession: body.FamProfession, FamPhone: body.FamPhone, FamAddress: body.FamAddress}
	result := db.DB.Create(&family)

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

func FamilyShow(c *gin.Context) {
	id := c.Param("id")

	var family models.Family
	err := db.DB.First(&family, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": family,
	})
}

func FamilyEdit(c *gin.Context) {
	id := c.Param("id")

	var body models.Family

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var family models.Family
	err := db.DB.First(&family, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&family).Updates(models.Family{
		FamName:       body.FamName,
		FamRelation:   body.FamRelation,
		FamProfession: body.FamProfession,
		FamPhone:      body.FamPhone,
		FamAddress:    body.FamAddress,
		EmployeeID:    body.EmployeeID,
		CandidateID:   body.CandidateID,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func FamilyDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var family models.Family
	err := db.DB.Delete(&family, "ID = ?", id).Error

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
