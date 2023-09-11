package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	GradeId  int    `gorm:"unique" json:"gradeid"`
	Min      int64  `json:"min"`
	Max      int64  `json:"max"`
	Struktur string `json:"struktur"`
}

type Increament struct {
	gorm.Model
	Percent string `json:"percent"`
}
