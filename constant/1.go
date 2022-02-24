package main

import (
	"fmt"
)

const (
	errorSpecialistType = iota
	catSepcialist
	dogSepcialist
	snakeSpecialist
)

func main() {
	var specialistType int
	fmt.Println(specialistType == errorSpecialistType)
	// fmt.Printf("%v", catSepcialist == specialistType)
}
