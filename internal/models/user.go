package models

import "gorm.io/gorm"

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	RoleID   uint
	Role     Role `gorm:"foreignKey:RoleID"`
}
