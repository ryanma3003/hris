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
		// &models.Grade{},
		// &models.Level{},
		// &models.JobDescription{},
		// &models.Division{},
		// &models.Department{},
		// &models.Supervision{},
		&models.Ptkp{},
		&models.Employee{},
		// &models.Candidate{},
		// &models.User{},
		// &models.Family{},
		// &models.Education{},
		// &models.Experience{},
		// &models.HealthDisease{},
		// &models.CriminalNote{},
		// &models.Course{},
		// &models.Reference{},
		// &models.SelfPerformance{},
		// &models.EvaluationForm{},
		// &models.EvaluationPoint{},
		// &models.Evaluation{},
		// &models.EvaluationResult{},
	)
}
