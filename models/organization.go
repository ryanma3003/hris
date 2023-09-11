package models

import "gorm.io/gorm"

type Division struct {
	gorm.Model
	Name string `json:"name"`
}

type Department struct {
	gorm.Model
	Name       string   `json:"name"`
	DivisionID int      `json:"divisionid"`
	Division   Division `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Position struct {
	gorm.Model
	Name         string     `json:"name"`
	DivisionID   int        `json:"divisionid"`
	Division     Division   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DepartmentID int        `json:"departmentid"`
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
