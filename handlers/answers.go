package handlers

import (
	"net/http"
	"qa-app/helpers"
	"qa-app/initializers"
	"qa-app/models"

	"github.com/gin-gonic/gin"
)

type VerifyRequest struct {
	QuestionID string `json:"question_id"`
	OptionID   string `json:"option_id"`
}

func VerifyAnswer(c *gin.Context) {
	var req VerifyRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the selected option
	var selectedOption models.Option
	result := initializers.DB.Where("option_id = ? AND question_id = ?", req.OptionID, req.QuestionID).First(&selectedOption)
	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusUnprocessableEntity, "Failed to verify answer", "01")
		return
	}

	// Check if the selected option is correct
	if selectedOption.IsCorrect {
		helpers.RespondWithSuccess(c, http.StatusOK, "Correct answer", "00")
	} else {
		helpers.RespondWithSuccess(c, http.StatusOK, "Wrong answer", "00")
	}
}
