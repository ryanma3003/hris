package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func PaidleaveIndex(c *gin.Context) {
	var pl []models.Paidleave
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

func PaidleaveCreate(c *gin.Context) {
	var body models.Paidleave

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse CutiReservedate dan CutiEnddate menjadi time.Time
	cutiReserveDate, err := time.Parse("2006-01-02", body.CutiReservedate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CutiReservedate format"})
		return
	}

	cutiEndDate, err := time.Parse("2006-01-02", body.CutiEnddate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CutiEnddate format"})
		return
	}

	// Hitung selisih hari
	days := int(cutiEndDate.Sub(cutiReserveDate).Hours() / 24)

	if body.CutiDays == nil {
		body.CutiDays = &days
	}

	// Inisialisasi nilai-nilai default
	defaultInteger := uint(0)
	defaultStatus := 0

	if body.Cutiverif1 == nil {
		body.Cutiverif1 = &defaultInteger
	}
	if body.Cutiverif2 == nil {
		body.Cutiverif2 = &defaultInteger
	}
	if body.Status == nil {
		body.Status = &defaultStatus
	}

	pl := models.Paidleave{
		EmployeeID:       body.EmployeeID,
		Reason:           body.Reason,
		CutiReservedate:  cutiReserveDate.Format("2006-01-02"),
		CutiEnddate:      cutiEndDate.Format("2006-01-02"),
		AlamatCuti:       body.AlamatCuti,
		PhoneCuti:        body.PhoneCuti,
		PenggantiCuti:    body.PenggantiCuti,
		PenggantiPhone:   body.PenggantiPhone,
		PenggantiJabatan: body.PenggantiJabatan,
		CutiDays:         body.CutiDays,
		Cutiverif1:       body.Cutiverif1,
		Cutiverif2:       body.Cutiverif2,
		Status:           body.Status,
	}
	result := db.DB.Create(&pl)

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

func PaidleaveShow(c *gin.Context) {
	id := c.Param("id")

	var pl models.Paidleave
	err := db.DB.First(&pl, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pl,
	})
}

func PaidleaveUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Paidleave

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var pl models.Paidleave
	err := db.DB.First(&pl, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	// Check if both Cutiverif1 and Cutiverif2 are equal to 1
	if body.Cutiverif1 != nil && body.Cutiverif2 != nil && *body.Cutiverif1 == 1 && *body.Cutiverif2 == 1 {
		// Set Status to 1
		pl.Status = new(int)
		*pl.Status = 1
	} else {
		// Set Status to something else if the condition is not met
		pl.Status = new(int)
		*pl.Status = 0 // or any other value you desire
	}

	// Update the Cutiverif1 and Cutiverif2 fields
	pl.Cutiverif1 = body.Cutiverif1
	pl.Cutiverif2 = body.Cutiverif2

	// Save the changes to the database
	db.DB.Save(&pl)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": pl,
	})
}

func PaidleaveDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var pl models.Paidleave
	err := db.DB.Delete(&pl, "ID = ?", id).Error

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

func PaidleaveCountStatus(c *gin.Context) {
	nikID := c.Param("nikid")

	// Menghitung total cuti masing-masing tahun
	type CutiPerTahun struct {
		Tahun                 int   `json:"tahun"`
		TotalCuti             int64 `json:"totalCuti"`
		TotalSisaCutiPerTahun int64 `json:"totalSisaCutiPerTahun"`
	}

	var cutiPerTahun []CutiPerTahun

	rows, err := db.DB.Model(&models.Paidleave{}).
		Select("EXTRACT(YEAR FROM TO_DATE(cuti_reservedate, 'DD Month YYYY')) as tahun, SUM(cuti_days) as total_cuti").
		Where("nik_id = ? AND status = ?", nikID, 1).
		Group("tahun").
		Order("tahun").
		Rows()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var totalCuti int64

	for rows.Next() {
		var tahun int
		var cutiTahunan int64
		rows.Scan(&tahun, &cutiTahunan)

		// Menghitung total cuti per tahun
		totalSisaCutiPerTahun := 12 - cutiTahunan
		if totalSisaCutiPerTahun < 0 {
			totalSisaCutiPerTahun = 0
		}

		totalCuti += cutiTahunan

		cutiPerTahun = append(cutiPerTahun, CutiPerTahun{
			Tahun:                 tahun,
			TotalCuti:             cutiTahunan,
			TotalSisaCutiPerTahun: totalSisaCutiPerTahun,
		})
	}

	c.JSON(http.StatusOK, cutiPerTahun)
}
