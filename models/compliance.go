package models

import (
	"gorm.io/gorm"
)

type SalarySlip struct {
	gorm.Model
	EmployeeID        uint     `json:"employeeid" gorm:"uniqueIndex:unique_employee_period"`
	Employee          Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete: NO ACTION;"`
	Period            string   `json:"period" gorm:"uniqueIndex:unique_employee_period"`
	Salary            int64    `json:"salary" gorm:"not null"`
	Status            int32    `json:"status"`
	SalarySlipDetails []SalarySlipDetail
}

type SalarySlipDetail struct {
	gorm.Model
	SalarySlipID uint       `json:"salaryslipid"`
	SalarySlip   SalarySlip `gorm:"foreignKey:SalarySlipID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Type         int32      `json:"type" gorm:"not null"`
	Name         string     `json:"name" gorm:"not null"`
	Value        int64      `json:"value" gorm:"not null"`
}

type Pph struct {
	gorm.Model
	Value      int64 `json:"value"`
	Percentage int32 `json:"percentage"`
}
