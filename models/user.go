package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	RoleID   uint   `json:"roleid"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// NikID    int64    `json:"nikid"`
	// NIK      Employee `gorm:"foreignKey:NikID;references:Nik"`
	EmployeeID uint     `json:"employeeid"`
	Employee   Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Password   string   `json:"-"`
}
