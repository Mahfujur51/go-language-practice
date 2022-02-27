//If ,Else If , Else With Input
package main

import "fmt"

func main() {
	var input int
	fmt.Println("Please Enter Number")
	fmt.Scanln(&input)
	fmt.Println(input)
	if input < 0 && input > 100 {
		fmt.Println("Please Enter Number Between 1 to 100")
	} else if input < 30 && input > 0 {
		fmt.Println("You are Fail")
	} else if input > 31 && input < 40 {
		fmt.Println("Your got D+")
	} else {
		fmt.Println("You are good Student")
	}
}
