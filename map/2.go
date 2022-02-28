package main

import "fmt"

func main() {
	mmap := map[int]string{
		20: "Geeks",
		30: "GFG",
		40: "Geeks for Geeks",
	}

	for i, v := range mmap {
		fmt.Println(i, "-", v)

	}
}
