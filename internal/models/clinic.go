package models

import "gorm.io/gorm"

type Clinic struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Address string
}
