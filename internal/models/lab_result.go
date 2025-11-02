package models

import "gorm.io/gorm"

type LabTestResult struct {
	gorm.Model
	LabTestOrderID uint
	LabTestOrder   LabTestOrder `gorm:"foreignKey:LabTestOrderID"`
	Result         string       `gorm:"type:text"`
	IsAbnormal     bool
}
