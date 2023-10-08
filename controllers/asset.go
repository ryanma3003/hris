package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func AssetIndex(c *gin.Context) {
	var pl []models.Asset
	if err := db.DB.Find(&pl).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pl,
	})
}

func CreateAsset(c *gin.Context) {
	var body models.Asset
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	as := models.Asset{
		EmployeeID:                body.EmployeeID,
		KategoriAset:              body.KategoriAset,
		KategoriAset_other:        body.KategoriAset_other,
		NamaAset:                  body.NamaAset,
		TahunAset:                 body.TahunAset,
		ModelAset:                 body.ModelAset,
		SerialNumberAset:          body.SerialNumberAset,
		StatusAset:                body.StatusAset,
		StatusAset_other:          body.StatusAset_other,
		KondisiAset:               body.KondisiAset,
		KondisiAset_other:         body.KondisiAset_other,
		ReasonKondisiAset:         body.ReasonKondisiAset,
		KondisiInternalAset:       body.KondisiInternalAset,
		ReasonKondisiInternalAset: body.ReasonKondisiInternalAset,
		PasanganAset:              body.PasanganAset,
		TahunPasanganAset:         body.TahunPasanganAset,
		KondisiPasanganAset:       body.KondisiPasanganAset,
	}
	result := db.DB.Create(&as)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// Return jumlah hari cuti
	c.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func AssetShow(c *gin.Context) {
	id := c.Param("id")

	var as models.Asset
	err := db.DB.First(&as, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": as,
	})
}

func AssetUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Asset

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var as models.Asset
	err := db.DB.First(&as, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&as).Updates(models.Asset{
		KategoriAset:              body.KategoriAset,
		KategoriAset_other:        body.KategoriAset_other,
		NamaAset:                  body.NamaAset,
		TahunAset:                 body.TahunAset,
		ModelAset:                 body.ModelAset,
		SerialNumberAset:          body.SerialNumberAset,
		StatusAset:                body.StatusAset,
		StatusAset_other:          body.StatusAset_other,
		KondisiAset:               body.KondisiAset,
		KondisiAset_other:         body.KondisiAset_other,
		ReasonKondisiAset:         body.ReasonKondisiAset,
		KondisiInternalAset:       body.KondisiInternalAset,
		ReasonKondisiInternalAset: body.ReasonKondisiInternalAset,
		PasanganAset:              body.PasanganAset,
		TahunPasanganAset:         body.TahunPasanganAset,
		KondisiPasanganAset:       body.KondisiPasanganAset,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": as,
	})
}

func AssetDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var as models.Asset
	err := db.DB.Delete(&as, "ID = ?", id).Error

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
