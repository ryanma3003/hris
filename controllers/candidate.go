package controllers

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"gorm.io/gorm"
)

func CandidateIndex(c *gin.Context) {
	var candidates []models.Candidate
	err := db.DB.Preload("JobDescription").Preload("Reqheadcount").Find(&candidates).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": candidates,
	})
}

func CandidateCreate(c *gin.Context) {
	// Get data req

	var body models.Candidate

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbFileName string
	if body.Avatar != "" {
		// decode
		b64data := body.Avatar[strings.IndexByte(body.Avatar, ',')+1:]
		imageData, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validation image file
		fileType := http.DetectContentType(imageData)
		if !strings.HasPrefix(fileType, "image/") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
			return
		}

		// extension & folder
		fileExtension := strings.TrimPrefix(fileType, "image/")
		uploadFolder := "uploads/"

		// check the folder if it exist
		if err := os.MkdirAll(uploadFolder, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileName := uuid.New().String() + "." + fileExtension
		dbFileName = uploadFolder + fileName

		// save image
		err = os.WriteFile(filepath.Join(uploadFolder, fileName), imageData, 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Create
	candidate := models.Candidate{
		Name:             body.Name,
		Avatar:           dbFileName,
		Email:            body.Email,
		JobDescriptionID: body.JobDescriptionID,
		ReqheadcountID:   body.ReqheadcountID,
		Type:             body.Type,
		ExpSalary:        body.ExpSalary,
		Address:          body.Address,
		Ciaddress:        body.Ciaddress,
		ExpBenefit:       body.ExpBenefit,
		Noktp:            body.Noktp,
		Npwp:             body.Npwp,
		Willing:          body.Willing,
		CompKnowledge:    body.CompKnowledge,
		WantJoin:         body.WantJoin,
		Phone:            body.Phone,
		Birthplace:       body.Birthplace,
		Birthdate:        body.Birthdate,
		Gender:           body.Gender,
		Religion:         body.Religion,
		Marital:          body.Marital,
		National:         body.National,
		Status:           body.Status,
	}

	result := db.DB.Create(&candidate)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func CandidateShow(c *gin.Context) {
	id := c.Param("id")

	var candidate models.Candidate
	err := db.DB.Preload("JobDescription").Preload("Reqheadcount").First(&candidate, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": candidate,
	})
}

func CandidateUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Candidate

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var candidate models.Candidate
	ok := db.DB.First(&candidate, "ID = ?", id).Error

	if ok != nil {
		errors.Is(ok, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	if body.Avatar != candidate.Avatar {
		// decode
		b64data := body.Avatar[strings.IndexByte(body.Avatar, ',')+1:]
		imageData, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// validation image file
		fileType := http.DetectContentType(imageData)
		if !strings.HasPrefix(fileType, "image/") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
			return
		}

		// extension & folder
		fileExtension := strings.TrimPrefix(fileType, "image/")
		uploadFolder := "uploads/"

		fileName := uuid.New().String() + "." + fileExtension

		// save image
		err = os.WriteFile(filepath.Join(uploadFolder, fileName), imageData, 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if candidate.Avatar != "" {
			// Define the path of the file to be deleted
			filePath := filepath.Join(candidate.Avatar)
			// Delete the file from the server
			err := os.Remove(filePath)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from upload folder"})
				return
			}
		}

		dbFileName := uploadFolder + fileName
		body.Avatar = dbFileName
	}

	// Update
	db.DB.Model(&candidate).Updates(models.Candidate{
		Name:             body.Name,
		Avatar:           body.Avatar,
		Email:            body.Email,
		JobDescriptionID: body.JobDescriptionID,
		ReqheadcountID:   body.ReqheadcountID,
		Type:             body.Type,
		ExpSalary:        body.ExpSalary,
		Address:          body.Address,
		Ciaddress:        body.Ciaddress,
		ExpBenefit:       body.ExpBenefit,
		Noktp:            body.Noktp,
		Npwp:             body.Npwp,
		Willing:          body.Willing,
		CompKnowledge:    body.CompKnowledge,
		WantJoin:         body.WantJoin,
		Phone:            body.Phone,
		Birthplace:       body.Birthplace,
		Birthdate:        body.Birthdate,
		Gender:           body.Gender,
		Religion:         body.Religion,
		Marital:          body.Marital,
		National:         body.National,
		Status:           body.Status,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func CandidateDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var candidate models.Candidate
	err := db.DB.Delete(&candidate, "ID = ?", id).Error

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
