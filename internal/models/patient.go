package models

import "gorm.io/gorm"

type PatientProfile struct {
	gorm.Model
	UserID        uint   `gorm:"unique;not null"`
	User          User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ContactEmail  string `gorm:"not null"`
	ContactPhone  string
	DateOfBirth   string
	MedicalHistory string `gorm:"type:text"`
}
