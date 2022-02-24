package main

import "fmt"

func main() {
	number := 50
	guess := 50
	if guess < 1 || guess >= 100 {
		fmt.Println("The guess Number must be between 1 to 100")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Number tow Low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it ")
		}
	}
}
