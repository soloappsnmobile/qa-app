package helpers

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Request Body
type Option struct {
	gorm.Model `json:"-"`
	OptionText string    `json:"option_text"`
	IsCorrect  bool      `json:"is_correct"`
	OptionID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"option_id"` // New UUID field
	QuestionID uint      `json:"-"`                                                     // This is the foreign key
}

type Question struct {
	gorm.Model `json:"-"`
	Question   string    `json:"question"`
	ID         uint      `json:"id" gorm:"primaryKey"`
	QuestionID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"question_id"`
	Options    []Option  `json:"options" gorm:"foreignKey:QuestionID"`
}

func ValidateQuestion(question *Question) error {
	if len(question.Options) < 2 {
		return errors.New("A question must have at least two options")
	}

	var correctAnswers []string

	for _, option := range question.Options {
		if option.IsCorrect {
			correctAnswers = append(correctAnswers, option.OptionText)
		}
	}

	if len(correctAnswers) == 0 {
		return errors.New("A question must have at least one correct answer")
	}

	if len(correctAnswers) > 1 {
		return errors.New("A question can only have one correct answer")
	}

	return nil
}

func ValidateOption(option *Option) error {
	if option.OptionText == "" {
		return errors.New("An option cannot be empty")
	}

	return nil
}
