package models

import "gorm.io/gorm"

type SelfPerformance struct {
	gorm.Model
	NikID       int64    `json:"nikid"`
	Employee    Employee `gorm:"foreignKey:NikID;references:Nik"`
	Listrespon  string   `json:"listrespone"`
	Perfrespon  string   `json:"perfrespone"`
	ListWorkobj string   `json:"listworkobj"`
	PerfWorkobj string   `json:"perfworkobj"`
	Corevalue   string   `json:"corevalue"`
	Perfvalue   string   `json:"perfvalue"`
	Team1       int64    `json:"team1"`
	Teamsatu    Employee `gorm:"foreignKey:Team1;references:Nik"`
	Team2       int64    `json:"team2"`
	Teamdua     Employee `gorm:"foreignKey:Team2;references:Nik"`
	Team3       int64    `json:"team3"`
	Teamtiga    Employee `gorm:"foreignKey:Team3;references:Nik"`
	Team4       int64    `json:"team4"`
	Teamempat   Employee `gorm:"foreignKey:Team4;references:Nik"`
}

type EvaluationForm struct {
	gorm.Model
	Name             string `json:"name"`
	EvaluationPoints []EvaluationPoint
}

type EvaluationPoint struct {
	gorm.Model
	EvaluationFormID int            `json:"evaluationformid"`
	EvaluationForm   EvaluationForm `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Point            int            `json:"point"`
	Desc             string         `json:"desc"`
}

type Evaluation struct {
	gorm.Model
	NikID             int64    `json:"nikid"`
	Employee          Employee `gorm:"foreignKey:NikID;references:Nik"`
	SubjectID         int64    `json:"subjectid"`
	Subject           Employee `gorm:"foreignKey:SubjectID;references:Nik"`
	Period            string   `json:"period"`
	EvaluationResults []EvaluationResult
}

type EvaluationResult struct {
	gorm.Model
	EvaluationID     int            `json:"evaluationid"`
	Evaluation       Evaluation     `gorm:"foreignKey:EvaluationID"`
	EvaluationFormID int            `json:"evaluationformid"`
	EvaluationForm   EvaluationForm `gorm:"foreignKey:EvaluationFormID"`
	Value            int            `json:"value"`
}
