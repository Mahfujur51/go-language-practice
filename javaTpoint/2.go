//Go Function with Multiple Return

package main

import "fmt"

func main() {
	fmt.Println(addAll(10, 20, 30, 40, 50))
}

func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0
	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value
	}
	return finalAddValue, finalSubValue
}
