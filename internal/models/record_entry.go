package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type RecordEntry struct {
	gorm.Model
	MedicalRecordID uint
	AuthorID        uint
	Author          User `gorm:"foreignKey:AuthorID"`
	Content         string `gorm:"type:text"`
	StructuredData  datatypes.JSON
}
