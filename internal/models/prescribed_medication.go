package models

import "gorm.io/gorm"

type PrescribedMedication struct {
	gorm.Model
	PrescriptionID uint
	Name           string
	Dosage         string
	Frequency      string
}
