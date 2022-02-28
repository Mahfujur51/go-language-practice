package main

import "fmt"

type Author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a Author) show() {
	fmt.Println("Author Name: ", a.name)
	fmt.Println("Author branch:", a.branch)
	fmt.Println("Publish Article: ", a.particles)
	fmt.Println("Author Salary:", a.salary)
}

func main() {

	a := Author{
		name:      "Mahfujur",
		branch:    "dhaka",
		particles: 300,
		salary:    20000,
	}
	a.show()

}
