package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Nik              int64          `json:"nik" gorm:"unique;primaryKey;autoIncrement:false"`
	Name             string         `json:"name"`
	Email            string         `json:"email" gorm:"unique"`
	GradeId          int            `json:"gradeid"`
	Grade            Grade          `gorm:"foreginKey:GradeId;references:Grade"`
	DivisionID       int            `json:"divisionid"`
	Division         Division       `gorm:"foreignKey:DivisionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DepartmentID     int            `json:"departmentid"`
	Department       Department     `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SupervisionID    int            `json:"supervisionid"`
	Supervision      Supervision    `gorm:"foreignKey:SupervisionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LevelID          int            `json:"levelid"`
	Level            Level          `gorm:"foreignKey:LevelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	JobDescriptionID int            `json:"jobdescriptionid"`
	JobDescription   JobDescription `gorm:"foreignKey:JobDescriptionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Salary           float64        `json:"salary"`
	Statusemployee   string         `json:"statusemployee"`
	Joindate         string         `json:"joindate"`
	Resigndate       string         `json:"resigndate"`
	Endcontract      string         `json:"endcontract"`
	Address          string         `json:"address"`
	Ciaddress        string         `json:"ciaddress"`
	Norek            string         `json:"norek"`
	Noktp            int64          `json:"noktp"`
	Npwp             string         `json:"npwp"`
	Kis              string         `json:"kis"`
	Kpj              string         `json:"kpj"`
	PtkpID           int            `json:"ptkpid"`
	Ptkp             Ptkp           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Phone            string         `json:"phone"`
	Birthplace       string         `json:"birthplace"`
	Birthdate        string         `json:"birthdate"`
	Gender           string         `json:"gender"`
	Religion         string         `json:"religion"`
	Marital          string         `json:"marital"`
	National         string         `json:"national"`
	Families         []Family
	Educations       []Education
	Experiences      []Experience
	HealthDiseases   []HealthDisease
	CriminalNotes    []CriminalNote
	Courses          []Course
	References       []Reference
}

type Ptkp struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique"`
	Value int    `json:"value"`
}

type Candidate struct {
	gorm.Model
	Name             string         `json:"nama"`
	JobDescriptionID int            `json:"jobdescriptionid"`
	JobDescription   JobDescription `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Type             string         `json:"type"`
	Email            string         `json:"email"`
	Phone            string         `json:"phone"`
	Gender           string         `json:"gender"`
	Birthplace       string         `json:"birthplace"`
	Birthdate        string         `json:"birthdate"`
	National         string         `json:"national"`
	Address          string         `json:"address"`
	Ciaddress        string         `json:"ciaddress"`
	Noktp            int64          `json:"noktp"`
	Npwp             string         `json:"npwp"`
	Religion         string         `json:"religion"`
	Marital          string         `json:"marital"`
	ExpSalary        float64        `json:"expsalary"`
	ExpBenefit       string         `json:"expbenefit"`
	Willing          string         `json:"willing"`
	CompKnowledge    string         `json:"compknowledge"`
	WantJoin         string         `json:"wantjoin"`
	Families         []Family
	Educations       []Education
	Experiences      []Experience
	HealthDiseases   []HealthDisease
	CriminalNotes    []CriminalNote
	Courses          []Course
	References       []Reference
}

type Family struct {
	gorm.Model
	EmployeeID    int       `json:"employeeid"`
	Employee      Employee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CandidateID   int       `json:"candidateid"`
	Candidate     Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FamRelation   string    `json:"famrelation"`
	FamName       string    `json:"famname"`
	FamProfession string    `json:"famprofession"`
	FamPhone      string    `json:"famphone"`
	FamAddress    string    `json:"famaddress"`
}

type Education struct {
	gorm.Model
	EmployeeID  int       `json:"employeeid"`
	Employee    Employee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CandidateID int       `json:"candidateid"`
	Candidate   Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Degree      string    `json:"degree"`
	YearComp    string    `json:"yearcomp"`
	Institute   string    `json:"institute"`
	Subject     string    `json:"subject"`
	GradePass   string    `json:"gradepass"`
}

type Experience struct {
	gorm.Model
	EmployeeID    int       `json:"employeeid"`
	Employee      Employee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CandidateID   int       `json:"candidateid"`
	Candidate     Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CompanyName   string    `json:"companyname"`
	CompJoinDate  string    `json:"compjoindate"`
	CompLeaveDate string    `json:"compleavedate"`
	PositionTitle string    `json:"positiontitle"`
	StatusEmp     string    `json:"statusemp"`
	Thp           float64   `json:"thp"`
	Gapok         float64   `json:"gapok"`
	Allowance     string    `json:"allowance"`
	OtherBenefit  string    `json:"otherbenefit"`
}

type HealthDisease struct {
	gorm.Model
	EmployeeID  int       `json:"employeeid"`
	Employee    Employee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CandidateID int       `json:"candidateid"`
	Candidate   Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CriminalNote struct {
	gorm.Model
	EmployeeID  int       `json:"employeeid"`
	Employee    Employee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CandidateID int       `json:"candidateid"`
	Candidate   Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Case        string    `json:"case"`
	Description string    `json:"description"`
}

type Course struct {
	gorm.Model
	EmployeeID  int       `json:"employeeid"`
	Employee    Employee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CandidateID int       `json:"candidateid"`
	Candidate   Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Institute   string    `json:"institute"`
	Type        string    `json:"type"`
	Venue       string    `json:"venue"`
	DateConduct string    `json:"dateconduct"`
	Competency  string    `json:"competency"`
}

type Reference struct {
	gorm.Model
	EmployeeID  int       `json:"employeeid"`
	Employee    Employee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CandidateID int       `json:"candidateid"`
	Candidate   Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RefName     string    `json:"refname"`
	RefPhone    string    `json:"refphone"`
	RefRelation string    `json:"refrelation"`
	RefTitle    string    `json:"reftitle"`
}
