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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func EmployeeIndex(c *gin.Context) {
	var employees []models.Employee
	err := db.DB.Preload("Division").Preload("Department").Preload("Supervision").Preload("JobDescription").Preload("Level").Preload("Grade").Preload("Ptkp").Find(&employees).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employees,
	})
}

func EmployeeCreate(c *gin.Context) {
	// Get data req

	var body models.Employee

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	employee := models.Employee{
		Name:             body.Name,
		Nik:              body.Nik,
		Email:            body.Email,
		GradeID:          body.GradeID,
		DivisionID:       body.DivisionID,
		DepartmentID:     body.DepartmentID,
		SupervisionID:    body.SupervisionID,
		LevelID:          body.LevelID,
		JobDescriptionID: body.JobDescriptionID,
		Salary:           body.Salary,
		Statusemployee:   body.Statusemployee,
		Joindate:         body.Joindate,
		Address:          body.Address,
		Ciaddress:        body.Ciaddress,
		Norek:            body.Norek,
		Noktp:            body.Noktp,
		Npwp:             body.Npwp,
		Kis:              body.Kis,
		Kpj:              body.Kpj,
		PtkpID:           body.PtkpID,
		Phone:            body.Phone,
		Birthplace:       body.Birthplace,
		Birthdate:        body.Birthdate,
		Gender:           body.Gender,
		Religion:         body.Religion,
		Marital:          body.Marital,
		National:         body.National,
	}

	result := db.DB.Create(&employee)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}

	// Hash the password
	dateString := body.Birthdate
	password := strings.Replace(dateString, "-", "", -1)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create user
	emailString := body.Email
	var username string
	at := strings.LastIndex(emailString, "@")
	if at >= 0 {
		username = emailString[:at]
	}

	user := models.User{Username: username, Password: string(hash), RoleID: 3, EmployeeID: employee.ID}
	resUser := db.DB.Create(&user)

	if resUser.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	message := fmt.Sprintf("Employee created successfully, and user login is generated automatically with \n username: %s \n password: %s", username, password)

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func EmployeeShow(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	err := db.DB.Preload("Division").Preload("Department").Preload("Supervision").Preload("JobDescription").Preload("Level").Preload("Grade").Preload("Ptkp").First(&employee, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}

func UpdateAvatar(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Find the data
	var employee models.Employee
	err := db.DB.First(&employee, "ID = ?", id).Error

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

		if employee.Avatar != "" {
			// Define the path of the file to be deleted
			filePath := filepath.Join(employee.Avatar)
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

	dbFileName = "uploads/" + newFileName

	// Update
	db.DB.Model(&employee).Updates(models.Employee{
		Avatar: dbFileName,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "avatar upload success",
	})
}

func GetAvatar(c *gin.Context) {
	id := c.Param("id")

	// Find the data
	var employee models.Employee
	err := db.DB.First(&employee, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	if employee.Avatar != "" {

		fileName := strings.Replace(employee.Avatar, "uploads/", "", 1)

		filePath := employee.Avatar
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

func EmployeeUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Employee

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var employee models.Employee
	err := db.DB.First(&employee, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&employee).Updates(models.Employee{
		Name:             body.Name,
		Nik:              body.Nik,
		Email:            body.Email,
		GradeID:          body.GradeID,
		DivisionID:       body.DivisionID,
		DepartmentID:     body.DepartmentID,
		SupervisionID:    body.SupervisionID,
		LevelID:          body.LevelID,
		JobDescriptionID: body.JobDescriptionID,
		Salary:           body.Salary,
		Statusemployee:   body.Statusemployee,
		Joindate:         body.Joindate,
		Address:          body.Address,
		Ciaddress:        body.Ciaddress,
		Norek:            body.Norek,
		Noktp:            body.Noktp,
		Npwp:             body.Npwp,
		Kis:              body.Kis,
		Kpj:              body.Kpj,
		PtkpID:           body.PtkpID,
		Phone:            body.Phone,
		Birthplace:       body.Birthplace,
		Birthdate:        body.Birthdate,
		Gender:           body.Gender,
		Religion:         body.Religion,
		Marital:          body.Marital,
		National:         body.National,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func EmployeeDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var employee models.Employee
	err := db.DB.Delete(&employee, "ID = ?", id).Error

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
