package models

import "gorm.io/gorm"

type JobDescription struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Employees   []Employee
}

type Level struct {
	gorm.Model
	Name      string `json:"name"`
	Employees []Employee
}

type Division struct {
	gorm.Model
	Name         string `json:"name"`
	Departments  []Department
	Supervisions []Supervision
	Employees    []Employee
}

type Department struct {
	gorm.Model
	Name         string   `json:"name"`
	DivisionID   int      `json:"divisionid"`
	Division     Division `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Supervisions []Supervision
	Employees    []Employee
}

type Supervision struct {
	gorm.Model
	Name         string     `json:"name"`
	DivisionID   int        `json:"divisionid"`
	Division     Division   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DepartmentID int        `json:"departmentid"`
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Employees    []Employee
}
