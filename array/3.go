package main

import "fmt"

func main() {
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
