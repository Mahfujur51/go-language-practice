package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}
type Employee struct {
	Person
	empid int
}

func (p Person) details() {
	fmt.Println(p, ""+"I am Person")
}

func (e Employee) details() {
	fmt.Println(e, ""+"I am employee")
}

func main() {
	x1 := Person{"Mahfujur", "Rahman"}
	x1.details()

	x2 := Employee{Person: Person{"Test", "Test rahman"}, empid: 11}
	x2.details()
}
