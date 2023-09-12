package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func EmployeeIndex(c *gin.Context) {
	var employees []models.Employee
	err := db.DB.Find(&employees).Error

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	employee := models.Employee{
		Name:             body.Name,
		Email:            body.Email,
		GradeId:          body.GradeId,
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
		Ptkp:             body.Ptkp,
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// Hash the password
	dateString := body.Birthdate.Format("01-02-2006")
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
	username := "string"
	at := strings.LastIndex(emailString, "@")
	if at >= 0 {
		username = emailString[:at]
	}

	user := models.User{Username: username, Password: string(hash), Role: "user"}
	resUser := db.DB.Create(&user)

	if resUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func EmployeeShow(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee
	err := db.DB.First(&employee, "ID = ?", id).Error

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

func EmployeeUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Employee

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
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
		Email:            body.Email,
		GradeId:          body.GradeId,
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
		Ptkp:             body.Ptkp,
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
		"data": employee,
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
