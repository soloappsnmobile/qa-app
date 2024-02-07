package main

import (
	"fmt"
	"log"
	"os"
	"qa-app/initializers"
	"qa-app/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := routers.SetupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router.Run(fmt.Sprintf(":%s", port))
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Create a list of questions
// 	questions := []string{
// 		"Who is the strongest Avenger?",
// 		"Who is the smartest Avenger?",
// 		"Who is the fastest Avenger?",
// 		"Who is the most handsome Avenger?",
// 		"Who is the most powerful Avenger?",
// 		"Who is the most beautiful Avenger?",
// 		"Who is the most intelligent Avenger?",
// 		"Who is the most charming Avenger?",
// 		"Who is the most attractive Avenger?",
// 		"Who is the most amazing Avenger?",
// 	}

// 	// Create a slice of options
// 	options := []string{
// 		"Thor",
// 		"Tony Stark",
// 		"Quicksilver",
// 		"Natasha Romanoff",
// 	}

// 	// Create a list of correct answers
// 	correctAnswers := []int{
// 		1, // Thor
// 		2, // Tony Stark
// 		3, // Quicksilver
// 		1, // Thor
// 		1, // Thor
// 		4, // Natasha Romanoff
// 		2, // Tony Stark
// 		2, // Tony Stark
// 		2, // Tony Stark
// 		2, // Tony Stark
// 	}

// 	fmt.Println("Welcome to the Avengers Quiz!")
// 	fmt.Println("Please answer the following questions:")
// 	fmt.Println()

// 	var totalScore, totalCorrectAnswers, totalWrongAnswers int

// 	for i, question := range questions {
// 		fmt.Println(question)
// 		for j, option := range options {
// 			fmt.Printf("%d: %s\n", j+1, option)
// 		}

// 		var answer int
// 		fmt.Scanln(&answer)

// 		if answer == correctAnswers[i] {
// 			fmt.Println("Correct!")
// 			totalScore++
// 			totalCorrectAnswers++
// 		} else {
// 			fmt.Println("Wrong!")
// 			totalWrongAnswers++
// 		}

// 		fmt.Println()
// 	}

// 	totalQuestions := len(questions)
// 	percentageCorrectAnswers := float64(totalCorrectAnswers) / float64(totalQuestions) * 100
// 	percentageWrongAnswers := float64(totalWrongAnswers) / float64(totalQuestions) * 100

// 	fmt.Println("Total Score:", totalScore)
// 	fmt.Println("Total Number of Questions:", totalQuestions)
// 	fmt.Println("Total Number of Correct Answers:", totalCorrectAnswers)
// 	fmt.Println("Total Number of Wrong Answers:", totalWrongAnswers)
// 	fmt.Println("Percentage of Correct Answers:", percentageCorrectAnswers, "%")
// 	fmt.Println("Percentage of Wrong Answers:", percentageWrongAnswers, "%")

// 	fmt.Println("Correct Answers:")
// 	for i, correctAnswer := range correctAnswers {
// 		fmt.Println(i+1, ":", options[correctAnswer-1])
// 	}
// }
