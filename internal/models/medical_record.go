package models

import "gorm.io/gorm"

type MedicalRecord struct {
	gorm.Model
	PatientID uint           `gorm:"unique;not null"`
	Patient   PatientProfile `gorm:"foreignKey:PatientID"`
	Entries   []RecordEntry  `gorm:"foreignKey:MedicalRecordID"`
}
