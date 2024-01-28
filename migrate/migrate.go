package main

import (
	"qa-app/initializers"
	"qa-app/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Question{})
	initializers.DB.AutoMigrate(&models.Option{})
}
