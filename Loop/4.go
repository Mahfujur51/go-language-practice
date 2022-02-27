//Goto in GO Language
package main

import "fmt"

func main() {
	var input int
Loop:
	fmt.Print("You are not Eligible to vote \n")
	fmt.Println("Enter your age")
	fmt.Scanln(&input)
	if input < 17 {
		goto Loop
	} else {
		fmt.Println("You can vote")
	}
}
