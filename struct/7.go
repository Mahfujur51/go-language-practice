package main

import "fmt"

type Student struct {
	name   string
	branch string
	year   int
}

type Teacher struct {
	name    string
	subject string
	exp     int
	details Student
}

func main() {
	result := Teacher{
		name:    "Mahfujur Rahman",
		subject: "CSE",
		exp:     5,
		details: Student{
			"Test Student", "Mirpur", 50,
		},
	}
	fmt.Println("Details Of Teacher:")
	fmt.Println("Teacher name: ", result.name)
	fmt.Println("Teacher subject:", result.subject)
	fmt.Println("Teacher Experience : ", result.exp)
	fmt.Println("******Details of Student :*****")
	fmt.Println("Student Name:", result.details.name)
	fmt.Println("Student Branch:", result.details.branch)
	fmt.Println("Student Year:", result.details.year)
	fmt.Println("Total Information:", result)
}
