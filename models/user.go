package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string   `gorm:"unique" json:"username"`
	Role       string   `json:"role"`
	NikID      int64    `json:"nikid"`
	NIK        Employee `gorm:"foreignKey:NikID;references:Nik"`
	EmployeeID int      `json:"employeeid"`
	Employee   Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Password   string   `json:"password"`
}
