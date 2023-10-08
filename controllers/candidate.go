package controllers

import (
	"errors"
	"fmt"
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
	err := db.DB.Find(&candidates).Error

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	candidate := models.Candidate{
		Name:             body.Name,
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

func UpdateAvatarCandidate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Find the data
	var candidate models.Candidate
	err := db.DB.First(&candidate, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	file, ok := c.FormFile("Avatar")
	var newFileName string
	var dbFileName string

	if ok == nil {

		if candidate.Avatar != "" {
			// Define the path of the file to be deleted
			filePath := filepath.Join(candidate.Avatar)
			// Delete the file from the server
			err := os.Remove(filePath)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file from upload folder"})
				return
			}
		}

		ext := filepath.Ext(file.Filename)
		newFileName = uuid.New().String() + ext
		if err := c.SaveUploadedFile(file, "uploads/"+newFileName); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}
	}

	dbFileName = "uploads/candidate/" + newFileName

	// Update
	db.DB.Model(&candidate).Updates(models.Employee{
		Avatar: dbFileName,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "avatar upload success",
	})
}

func GetAvatarCandidate(c *gin.Context) {
	id := c.Param("id")

	// Find the data
	var candidate models.Candidate
	err := db.DB.First(&candidate, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	if candidate.Avatar != "" {

		fileName := strings.Replace(candidate.Avatar, "uploads/", "", 1)

		filePath := candidate.Avatar
		// Open the file
		fileData, err := os.Open(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer fileData.Close()
		// Read the first 512 bytes of the file to determine its content type
		fileHeader := make([]byte, 512)
		_, err = fileData.Read(fileHeader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}
		fileContentType := http.DetectContentType(fileHeader)
		// Get the file info
		fileInfo, err := fileData.Stat()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
			return
		}
		// Set the headers for the file transfer and return the file
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		c.Header("Content-Type", fileContentType)
		c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
		c.File(filePath)
	}
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
	err := db.DB.First(&candidate, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&candidate).Updates(models.Candidate{
		Name:             body.Name,
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
