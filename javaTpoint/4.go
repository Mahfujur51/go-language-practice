// Struct in Go Language

package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	x := Person{"Mahfujur", "Rahman", 9}
	fmt.Println(x)
	fmt.Println(x.firstName)

}
