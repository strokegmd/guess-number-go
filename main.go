package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func game(chances int) {
	var guess int

	rand.Seed(uint64(time.Now().UnixNano()))
	number := rand.Intn(100)

	fmt.Print("Let's start the game!\n")
	for attempts := 0; chances > 0; chances-- {
		fmt.Print("\nEnter your guess: ")
		fmt.Scan(&guess)

		if number > guess {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		} else if number < guess {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			break
		}

		attempts++
	}
}

func main() {
	var choice int
	var chances int

	fmt.Print(
		"Welcome to the Number Guessing Game!\n",
		"I'm thinking of a number between 1 an 100.\n",
		"You have 5 chances to guess the correct number.\n\n",
		"Please select the difficulty level:\n",
		"1. Easy (10 chances)\n",
		"2. Medium (5 chances)\n",
		"3. Hard (3 chances)\n\n",
		"Enter your choice: ",
	)

	for {
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("\nGreat! You have selected the Easy difficulty level.")
			chances = 10
		case 2:
			fmt.Println("\nGreat! You have selected the Medium difficulty level.")
			chances = 5
		case 3:
			fmt.Println("\nGreat! You have selected the Hard difficulty level.")
			chances = 3
		default:
			fmt.Print("\nIncorrect choice. Enter your choice: ")
		}

		if chances != 0 {
			game(chances)
			break
		}
	}
}
