package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I,Love,My,Country"
	var arr []string = strings.Split(str, ",")
	fmt.Println(len(arr))
	for index, value := range arr {
		fmt.Println("Index : ", index, "value : ", value)
		// fmt.Println(value)
	}
}
