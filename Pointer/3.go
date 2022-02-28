package main

import "fmt"

func main() {
	ptr := new(int)
	fmt.Println("berfor Change", ptr)
	chagePtr(ptr)

}

func chagePtr(ptr *int) {
	*ptr = 0
}
