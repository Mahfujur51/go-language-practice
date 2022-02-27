//Function and Struct ;
package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
}

func (emp Employee) fullName() {
	fmt.Println(emp.firstName + " " + emp.lastName)
}
func main() {
	e1 := Employee{"Mahfujur", "Rahman"}
	e1.fullName()

}
