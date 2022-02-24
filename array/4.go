package main

import "fmt"

func main() {
	a := []int{}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, []int{2, 3, 4}...)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}
