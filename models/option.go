package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Option struct {
	OptionID   uuid.UUID      `json:"option_id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	OptionText string         `json:"option_text"`
	IsCorrect  bool           `json:"-"`
	QuestionID uuid.UUID      `json:"-" gorm:"type:uuid"` // This is the foreign key
}
