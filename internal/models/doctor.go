package models

import "gorm.io/gorm"

type DoctorProfile struct {
	gorm.Model
	UserID      uint
	User        User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpecialtyID uint
	Specialty   Specialty `gorm:"foreignKey:SpecialtyID"`
}
