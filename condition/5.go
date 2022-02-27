package main

import "fmt"

func main() {
	var Input int
	fmt.Print("Please Inter Number : ")
	fmt.Scanln(&Input)
	switch Input {
	case 10:
		fmt.Println("Your Number is 10")
	case 20:
		fmt.Println("Your Number is 20")
	default:
		fmt.Println("You Enter Wrong Number")

	}
}
