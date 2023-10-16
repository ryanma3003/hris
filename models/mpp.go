package models

import "gorm.io/gorm"

type Mpp struct {
	gorm.Model
	EmployeeID    uint           `json:"employeeid"`
	Employee      Employee       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	Period        string         `json:"period"`
	DivisionID    uint           `json:"divisionid"`
	Division      Division       `gorm:"foreignKey:DivisionID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	Numberreq     int            `json:"numberreq"`
	Budget        int            `json:"budget"`
	Status        int            `json:"status"`
	Reqheadcounts []Reqheadcount `gorm:"foreignKey:MppID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Reqheadcount struct {
	gorm.Model
	MppID            uint           `json:"mppid"`
	Mpp              Mpp            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION"`
	EmployeeID       uint           `json:"employeeid"`
	Employee         Employee       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	LevelID          int            `json:"levelid"`
	Level            Level          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	GradeID          int            `json:"gradeid"`
	Grade            Grade          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	Statusemployee   string         `json:"statusemployee"`
	Reasonhiring     string         `json:"reasonhiring"`
	Degree           string         `json:"degree"`
	Minexp           string         `json:"minexp"`
	JobDescriptionID uint           `json:"jobdescriptionid"`
	JobDescription   JobDescription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	Specification    string         `json:"specification"`
	Gender           string         `json:"gender"`
	Age              string         `json:"age"`
	Maritalstatus    string         `json:"maritalstatus"`
	Recruitmenttype  string         `json:"recruitmenttype"`
	Status           string         `json:"status"`
}
