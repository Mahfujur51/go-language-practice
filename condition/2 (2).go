//If Else with form input
package main

import "fmt"

func main() {

	fmt.Println("Enter Number")
	var input int
	fmt.Scanln(&input)
	fmt.Print(input)
	if input%2 == 0 {
		fmt.Printf("Number is Even")
	} else {
		fmt.Printf("Number is Odd")
	}
}
