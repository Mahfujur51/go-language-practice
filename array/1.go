package main

import "fmt"

func main() {
	// grades := [3]int{10, 30, 40}
	// fmt.Println(grades)
	a := [...]int{10, 20, 30, 50}
	b := a
	b[1] = 50
	fmt.Println(a)
	fmt.Println(len(b))
	fmt.Println(cap(a))
}
