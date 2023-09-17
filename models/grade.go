package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Grade     int        `gorm:"unique" json:"grade"`
	Min       int64      `json:"min"`
	Max       int64      `json:"max"`
	Struktur  string     `json:"struktur"`
	Employees []Employee `gorm:"foreignKey:GradeId;references:Grade"`
}

type Increament struct {
	gorm.Model
	Percent string `json:"percent"`
}
