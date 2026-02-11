package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 to 100")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess) 

		if guess == target {
			fmt.Println("Congratulations you guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}