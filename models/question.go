package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model `json:"-"`
	Question   string   `json:"question"`
	ID         uint     `json:"id" gorm:"primaryKey"`
	Options    []Option `json:"options" gorm:"foreignKey:QuestionID"`
}

// gorm:"foreignKey:QuestionID"
