package models

import "gorm.io/gorm"

type NurseProfile struct {
	gorm.Model
	UserID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
