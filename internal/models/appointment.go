package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PatientID uint
	Patient   PatientProfile `gorm:"foreignKey:PatientID"`
	DoctorID  uint
	Doctor    DoctorProfile `gorm:"foreignKey:DoctorID"`
	ClinicID  uint
	Clinic    Clinic `gorm:"foreignKey:ClinicID"`
	Date      time.Time
	Status    string // e.g., Scheduled, Canceled, Completed
}
