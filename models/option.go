package models

import "gorm.io/gorm"

// type Option struct {
// 	gorm.Model `json:"-"`
// 	Option     string `json:"option_text"`
// 	ID         uint   `json:"id" gorm:"primaryKey"`
// 	IsCorrect  bool   `json:"is_correct"`
// }

type Option struct {
	gorm.Model `json:"-"`
	OptionText string `json:"option_text"`
	IsCorrect  bool   `json:"is_correct"`
	QuestionID uint   `json:"-"` // This is the foreign key
}
