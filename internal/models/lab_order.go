package models

import "gorm.io/gorm"

type LabTestOrder struct {
	gorm.Model
	PatientID uint
	Patient   PatientProfile `gorm:"foreignKey:PatientID"`
	DoctorID  uint
	Doctor    DoctorProfile `gorm:"foreignKey:DoctorID"`
	TestName  string
	Status    string // e.g., Ordered, Completed
}
