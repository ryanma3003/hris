package main

import (
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(
		&models.Grade{},
		&models.JobDescription{},
		&models.Level{},
		&models.Division{},
		&models.Department{},
		&models.Supervision{},
		&models.Employee{},
		&models.User{},
	)
}
