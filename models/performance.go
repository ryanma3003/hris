package models

import "gorm.io/gorm"

type SelfPerformance struct {
	gorm.Model
	EmployeeID  uint     `json:"employeeid"`
	Employee    Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete: NO ACTION;"`
	Listrespon  string   `json:"listrespone"`
	Perfrespon  string   `json:"perfrespone"`
	ListWorkobj string   `json:"listworkobj"`
	PerfWorkobj string   `json:"perfworkobj"`
	Corevalue   string   `json:"corevalue"`
	Perfvalue   string   `json:"perfvalue"`
	Team1       uint     `json:"team1"`
	Teamsatu    Employee `gorm:"foreignKey:Team1"`
	Team2       uint     `json:"team2"`
	Teamdua     Employee `gorm:"foreignKey:Team2"`
	Team3       uint     `json:"team3"`
	Teamtiga    Employee `gorm:"foreignKey:Team3"`
	Team4       uint     `json:"team4"`
	Teamempat   Employee `gorm:"foreignKey:Team4"`
}

type EvaluationTemplate struct {
	gorm.Model
	Name            string            `json:"name"`
	EvaluationForms []*EvaluationForm `gorm:"many2many:evtemplates_evforms;"`
}

type EvaluationForm struct {
	gorm.Model
	Name                string `json:"name"`
	EvaluationPoints    []EvaluationPoint
	EvaluationTemplates []*EvaluationTemplate `gorm:"many2many:evtemplates_evforms;"`
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
	EmployeeID        uint     `json:"employeeid"`
	Employee          Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete: NO ACTION;"`
	SubjectID         uint     `json:"subjectid"`
	Subject           Employee `gorm:"foreignKey:SubjectID"`
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
