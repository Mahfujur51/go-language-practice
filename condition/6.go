package main

import "fmt"

func main() {
	k := 30
	switch k {
	case 10:
		fmt.Println("Number is 10")
		fallthrough
	case 20:
		fmt.Println("Your Number is 30")
		fallthrough
	case 30:
		fmt.Println("Your Number is 30")
		fallthrough
	default:
		fmt.Println("Test")

	}
}
