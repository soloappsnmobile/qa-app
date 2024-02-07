package routers

import (
	"qa-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/questions", handlers.CreateQuestion)
	router.GET("/questions", handlers.GetQuestions)
	router.DELETE("/questions/:id", handlers.DeleteQuestion)
	router.GET("/questions/:id", handlers.GetQuestion)
	router.PUT("/questions/:id", handlers.UpdateQuestion)

	router.POST("/verify-answer", handlers.VerifyAnswer)

	return router
}
