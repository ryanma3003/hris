package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func MppIndex(c *gin.Context) {
	division := c.Query("division")

	var mpps []models.Mpp
	err := db.DB
	if division != "" {
		err = err.Select("DISTINCT ON (division_id) *").Preload("Division").Find(&mpps, "division_id = ?", division)
	} else {
		err = err.Select("DISTINCT ON (division_id) *").Preload("Division").Find(&mpps)
	}

	if err.Error != nil {
		errors.Is(err.Error, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mpps,
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func MppListUnapprove(c *gin.Context) {
	var mppData []models.Mpp

	err := db.DB.Find(&mppData, "status = ?", 0).Error
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mppData,
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func MppCreate(c *gin.Context) {
	// Get data req
	var body []models.Mpp

	fmt.Printf("%v", body)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	empID, _ := c.Get("id")
	divID, _ := c.Get("division")

	for i := range body {
		mpp := models.Mpp{
			EmployeeID: empID.(uint),
			Period:     body[i].Period,
			DivisionID: divID.(uint),
			Numberreq:  body[i].Numberreq,
			Budget:     body[i].Budget,
			Status:     body[i].Status,
		}

		result := db.DB.Create(&mpp)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": result.Error,
			})
			return
		}
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "Created Success",
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func MppUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data req
	var body models.Mpp

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var mppData models.Mpp
	err := db.DB.First(&mppData, "id = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	err = db.DB.Model(&mppData).Updates(models.Mpp{
		Period:     body.Period,
		DivisionID: body.DivisionID,
		Numberreq:  body.Numberreq,
		Budget:     body.Budget,
	}).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated Success",
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func MppShow(c *gin.Context) {
	id := c.Param("id")

	var mppData models.Mpp
	err := db.DB.First(&mppData, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mppData,
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func ApproveMpp(c *gin.Context) {
	id := c.Param("id")

	var mppData models.Mpp
	err := db.DB.First(&mppData, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	err = db.DB.Model(&mppData).Updates(models.Mpp{Status: 2}).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Approved Success",
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func RevisionMpp(c *gin.Context) {
	id := c.Param("id")

	var mppData models.Mpp
	err := db.DB.First(&mppData, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	err = db.DB.Model(&mppData).Updates(models.Mpp{Status: 1}).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Revision Success",
	})
}
