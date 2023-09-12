package models

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
	Name string `json:"name"`
}

type Codedept struct {
	gorm.Model
	Name string `json:"name"`
}
