package handlers

import (
	"net/http"
	"qa-app/helpers"
	"qa-app/initializers"
	"qa-app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateQuestion creates a question in the database
func CreateQuestion(c *gin.Context) {
	var question helpers.Question

	if err := c.Bind(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := helpers.ValidateQuestion(&question)
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "01")
		return
	}

	questionModel := models.Question{
		Question: question.Question,
	}

	for _, option := range question.Options {
		questionModel.Options = append(questionModel.Options, models.Option{
			OptionText: option.OptionText,
			IsCorrect:  option.IsCorrect,
		})
	}

	result := initializers.DB.Create(&questionModel)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Failed to create question", "01")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusCreated, "Question created successfully", "00", questionModel)
}

// GetQuestions retrieves all questions from the database
func GetQuestions(c *gin.Context) {
	var questions []models.Question

	// Use GORM's Find() method to retrieve all questions from the database. When getting it, query the options as well
	result := initializers.DB.Preload("Options").Find(&questions)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error.Error(), "01")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "Questions retrieved successfully", "00", questions)
}

// GetQuestion retrieves a single question from the database
func GetQuestion(c *gin.Context) {
	var question models.Question
	questionID := c.Param("id")

	// Use GORM's First() method to retrieve the question from the database. When getting it, query the options as well
	result := initializers.DB.Preload("Options").Where("id = ?", questionID).First(&question)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error.Error(), "01")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "Question retrieved successfully", "00", question)
}

// UpdateQuestion updates a question in the database
func UpdateQuestion(c *gin.Context) {
	var question helpers.Question
	questionID := c.Param("id")

	if err := c.Bind(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := helpers.ValidateQuestion(&question)
	if err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "01")
		return
	}

	// Update the question
	questionModel := models.Question{
		Question: question.Question,
	}

	result := initializers.DB.Model(&models.Question{}).Where("id = ?", questionID).Updates(&questionModel)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Failed to update question", "01")
		return
	}

	// Delete the old options
	var options []models.Option
	result = initializers.DB.Where("question_id = ?", questionID).Delete(&options)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, result.Error.Error(), "01")
		return
	}

	// Create new options
	for _, option := range question.Options {
		newOption := models.Option{
			OptionText: option.OptionText,
			IsCorrect:  option.IsCorrect,
			QuestionID: uuid.MustParse(questionID), // Make sure to set the QuestionID
		}
		result = initializers.DB.Create(&newOption)
		if result.Error != nil {
			helpers.RespondWithError(c, http.StatusInternalServerError, result.Error.Error(), "01")
			return
		}
		questionModel.Options = append(questionModel.Options, newOption)
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "Question updated successfully", "00", questionModel)
}

// DeleteQuestion deletes a question from the database
func DeleteQuestion(c *gin.Context) {
	var question models.Question
	questionID := c.Param("id")

	// Delete the options associated with the question
	var options []models.Option
	result := initializers.DB.Where("question_id = ?", questionID).Delete(&options)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, result.Error.Error(), "01")
		return
	}

	// Delete the question
	result = initializers.DB.Where("id = ?", questionID).Delete(&question)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, result.Error.Error(), "01")
		return
	}

	helpers.RespondWithSuccess(c, http.StatusOK, "Question deleted successfully", "00")
}
