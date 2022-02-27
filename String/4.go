//String Splite
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I,Love,My,Country"
	var arr []string = strings.Split(str, ",")
	fmt.Println(len(arr))
	for i, v := range arr {
		fmt.Println("Index:", i, "Value:", v)
	}
}
