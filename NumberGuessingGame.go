package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates and returns a random number between 1 and 100.
func getRandomNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(100) + 1
}

// Displays difficulty options, reads user choice, and returns allowed chances.
func selectDifficulty() int {
	for {
		fmt.Print("\nPlease select the difficulty level:")
		fmt.Println("\n1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")

		var level int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanln(&level)
		if err != nil {
			fmt.Println("Enter a numeric choice 1, 2 or 3.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch level {
		case 1:
			fmt.Println("Great! You have selected the Easy difficulty level.")
			return 10
		case 2:
			fmt.Println("Great! You have selected the Medium difficulty level.")
			return 5
		case 3:
			fmt.Println("Great! You have selected the Hard difficulty level.")
			return 3
		default:
			fmt.Println("Invalid Selection. Select only 1, 2 or 3.")

		}
	}
}

// Runs one complete game round including guessing logic and win/lose handling.
func playRound(chances int) {
	generatedNumber := getRandomNumber()

	var attempts int

	for chances > 0 {
		var guessedNumber int
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanln(&guessedNumber)

		if err != nil {
			fmt.Println("Invalid input! Please enter a whole number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if guessedNumber > 100 || guessedNumber < 1 {
			fmt.Println("Enter a number between 1 and 100.")
			continue
		}

		attempts++

		if guessedNumber > generatedNumber {
			fmt.Println("Incorrect! The number is less than", guessedNumber)
			chances--
			fmt.Printf("Remaining chances: %d\n", chances)

		} else if guessedNumber < generatedNumber {
			fmt.Println("Incorrect! The number is greater than", guessedNumber)
			chances--
			fmt.Printf("Remaining chances: %d\n", chances)

		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			break
		}

	}

	if chances == 0 {
		fmt.Printf("Game Over! You've used all your chances. The correct number was %d\n", generatedNumber)
	}
}

// Controls overall program flow including replay loop.
func main() {
	for {
		fmt.Println("Welcome to the Number Guessing Game!")
		fmt.Println("I'll let you choose a number between 1 and 100.")

		chances := selectDifficulty()

		if chances == 0 {
			continue
		}

		fmt.Printf("You have %d chances to guess the correct number.\n", chances)

		fmt.Println("Let's start the game!")

		playRound(chances)

		fmt.Print("Do you want to play again? (y/n): ")

		var playAgain string
		fmt.Scanln(&playAgain)

		if len(playAgain) == 0 || (playAgain[0] != 'y' && playAgain[0] != 'Y') {
			fmt.Println("Thanks for playing!")
			break
		}

		fmt.Println("----------------------------")
		fmt.Println("Starting new game...")
		fmt.Println("----------------------------")
	}

}
