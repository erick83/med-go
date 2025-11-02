package models

import "gorm.io/gorm"

type Specialty struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
