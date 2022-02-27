package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6, 10}
	sum := 0

	for _, value := range nums {
		// fmt.Println(y, "-", value)
		sum += value

	}
	fmt.Println("sum: ", sum)

}
