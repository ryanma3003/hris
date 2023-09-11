package models

import "gorm.io/gorm"

type Month struct {
	gorm.Model
	Name  string `json:"name"`
	Sname string `json:"sname"`
}
