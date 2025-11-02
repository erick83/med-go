package models

import "gorm.io/gorm"

type Prescription struct {
	gorm.Model
	PatientID uint
	Patient   PatientProfile `gorm:"foreignKey:PatientID"`
	DoctorID  uint
	Doctor    DoctorProfile `gorm:"foreignKey:DoctorID"`
	Notes     string
	Medications []PrescribedMedication `gorm:"foreignKey:PrescriptionID"`
}
