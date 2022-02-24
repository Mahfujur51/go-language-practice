package main

import (
	"fmt"
)

func main() {
	setPopulation := make(map[string]int)
	setPopulation = map[string]int{
		"Dhaka":    1534566,
		"Khulna":   564412,
		"Sylhet":   357454,
		"Jashore":  575454,
		"Meherpur": 6785466874,
	}
	setPopulation["Rajshahi"] = 300000
	fmt.Println(setPopulation)
}
