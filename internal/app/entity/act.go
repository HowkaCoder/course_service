package entity

import "gorm.io/gorm"

type Act struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	AnswerID uint   `gorm:"not null"`
	ActText  string `gorm:"type:varchar(500);not null"`
	ImageUrl string `gorm:"null"`
	Position uint   `gorm:"not null"`
}
