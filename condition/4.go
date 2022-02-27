// Nestend If Else
package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20

	if x > 0 {
		fmt.Println("You are in X condition Block")
		if y == 20 {
			fmt.Println("Your Number is 20")
		}

	}
}
