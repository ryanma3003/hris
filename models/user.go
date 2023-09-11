package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string   `gorm:"unique" json:"username"`
	EmployeeId string   `json:"employeeid"`
	Employee   Employee `gorm:"references:EmployeeId"`
	Password   string   `json:"password"`
}
