package models

import "gorm.io/gorm"

type Paidleave struct {
	gorm.Model
	EmployeeID       uint     `json:"employeeid"`
	Employee         Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reason           string   `json:"reason"`
	CutiReservedate  string   `json:"cutireservedate"`
	CutiEnddate      string   `json:"cutienddate"`
	AlamatCuti       string   `json:"alamatcuti"`
	PhoneCuti        string   `json:"phonecuti"`
	PenggantiCuti    int      `json:"pengganticuti"`
	PenggantiPhone   string   `json:"penggantiphone"`
	PenggantiJabatan string   `json:"penggantijabatan"`
	CutiDays         *int     `json:"cutidays"`
	Cutiverif1       *uint    `json:"cutiverif1"`
	Cutiverifsatu    Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cutiverif2       *uint    `json:"cutiverif2"`
	Cutiverifdua     Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status           *int     `json:"status"`
}

type Asset struct {
	gorm.Model
	EmployeeID                uint     `json:"employeeid"`
	Employee                  Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	KategoriAset              string
	KategoriAset_other        string
	NamaAset                  string
	TahunAset                 int
	ModelAset                 string
	SerialNumberAset          string
	StatusAset                string
	StatusAset_other          string
	KondisiAset               string
	KondisiAset_other         string
	ReasonKondisiAset         string
	KondisiInternalAset       string
	ReasonKondisiInternalAset string
	PasanganAset              string
	TahunPasanganAset         string
	KondisiPasanganAset       string
}

type Loan struct {
	gorm.Model
	LoanID     string   `json:"loanid"`
	EmployeeID uint     `json:"employeeid"`
	Employee   Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reason     string   `json:"reason"`
	Amount     string   `json:"amount"`
}

type Insurance struct {
	gorm.Model
	EmployeeID uint     `json:"employeeid"`
	Employee   Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Reason     string   `json:"reason"`
	Verified1  *int     `json:"verified1"`
	Verified2  *int     `json:"verified2"`
	Status     *int     `json:"status"`
}
