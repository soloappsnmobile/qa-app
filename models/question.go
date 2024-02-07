package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model `json:"-"`
	Question   string    `json:"question"`
	ID         uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Options    []Option  `json:"options" gorm:"foreignKey:QuestionID"`
}

// gorm:"foreignKey:QuestionID"
