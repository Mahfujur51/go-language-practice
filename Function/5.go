package main

import "fmt"

func area(a int, b int) int {
	arr := a * b
	return arr
}

func main() {
	x := area(20, 30)

	fmt.Println("The area ", x)

}
