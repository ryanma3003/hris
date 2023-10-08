package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func MppIndex(c *gin.Context) {
	employeeid := c.Param("employeeid")

	var mpps []models.Mpp
	err := db.DB.Find(&mpps, "employee_id = ?", employeeid).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
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
func MppCreate(c *gin.Context) {
	// Get data req
	var body []models.Mpp

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range body {
		mpp := models.Mpp{
			EmployeeID: body[i].EmployeeID,
			Period:     body[i].Period,
			DivisionID: body[i].DivisionID,
			Numberreq:  body[i].Numberreq,
			Budget:     body[i].Budget,
			Status:     body[i].Status,
		}

		result := db.DB.Create(&mpp)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": result.Error,
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
		"data": "Updated Success",
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
		"data": "Approved Success",
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
		"data": "Revision Success",
	})
}
