package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
}