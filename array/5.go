package main

import "fmt"

func main() {
	myarray := [...]int{100, 200, 300, 400, 500, 600, 700}
	fmt.Println("Original Array (Before):", myarray)
	newArray := myarray
	fmt.Println("New Array", newArray)
	newArray[0] = 1000
	fmt.Println("Modified Array", newArray)
	fmt.Println("Original Array", myarray)
}
