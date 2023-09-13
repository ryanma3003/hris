package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string   `gorm:"unique" json:"username"`
	Role       string   `json:"role"`
	EmployeeId int64    `json:"employeeid"`
	Employee   Employee `gorm:"references:EmployeeId"`
	Password   string   `json:"password"`
}
