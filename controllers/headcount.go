package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func ListMpp(c *gin.Context) {
	employeeid := c.Param("employeeid")
	period := c.Param("period")

	var mpps []models.Mpp
	err := db.DB.Where("employee_id = ? AND period = ?", employeeid, period).Find(&mpps).Error

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
func FormHeadcount(c *gin.Context) {
	mppid := c.Param("mppid")

	var mpp models.Mpp
	err := db.DB.First(&mpp, "ID = ?", mppid).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": mpp,
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func CreateHeadcount(c *gin.Context) {
	// Get data req
	var body models.Reqheadcount

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi jumlah pengajuan (based on MPP)
	var mppData models.Mpp
	err := db.DB.First(&mppData, "id = ?", body.MppID).Error
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	var data models.Reqheadcount
	var count int64
	db.DB.Where("mpp_id = ?", body.MppID).Find(&data).Count(&count)

	if count >= int64(mppData.Numberreq) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Number of Request beyond annual MPP, can't apply again"})
		return
	}

	// Insert
	headcount := models.Reqheadcount{
		MppID:            body.MppID,
		EmployeeID:       body.EmployeeID,
		LevelID:          body.LevelID,
		GradeID:          body.GradeID,
		Statusemployee:   body.Statusemployee,
		Reasonhiring:     body.Reasonhiring,
		Degree:           body.Degree,
		Minexp:           body.Minexp,
		JobDescriptionID: body.JobDescriptionID,
		Specification:    body.Specification,
		Gender:           body.Gender,
		Age:              body.Age,
		Maritalstatus:    body.Maritalstatus,
		Recruitmenttype:  body.Recruitmenttype,
		Status:           body.Status,
	}

	result := db.DB.Create(&headcount)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": result.Error,
		})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "Created Success",
		// "message": count,
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func ShowAllHeadcount(c *gin.Context) {
	employeeid := c.Param("employeeid")

	var headcounts []models.Reqheadcount
	err := db.DB.Where("employee_id = ?", employeeid).Find(&headcounts).Error
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": headcounts,
	})
}

// ======================================================================================================
// ======================================================================================================
// ======================================================================================================
func UpdateHeadcount(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data req
	var body models.Reqheadcount

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var headcountData models.Reqheadcount
	err := db.DB.First(&headcountData, "id = ?", id).Error
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	err = db.DB.Model(&headcountData).Updates(models.Reqheadcount{
		LevelID:          body.LevelID,
		GradeID:          body.GradeID,
		Statusemployee:   body.Statusemployee,
		Reasonhiring:     body.Reasonhiring,
		Degree:           body.Degree,
		Minexp:           body.Minexp,
		JobDescriptionID: body.JobDescriptionID,
		Specification:    body.Specification,
		Gender:           body.Gender,
		Age:              body.Age,
		Maritalstatus:    body.Maritalstatus,
		Recruitmenttype:  body.Recruitmenttype,
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
func ApproveReqHeadcount(c *gin.Context) {
	id := c.Param("id")

	var headcountData models.Reqheadcount
	err := db.DB.First(&headcountData, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	err = db.DB.Model(&headcountData).Updates(models.Reqheadcount{Status: "2"}).Error

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
func RevisionReqHeadcount(c *gin.Context) {
	id := c.Param("id")

	var headcountData models.Reqheadcount
	err := db.DB.First(&headcountData, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	err = db.DB.Model(&headcountData).Updates(models.Reqheadcount{Status: "1"}).Error

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
