package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:     3,
		actorName:  "Mahfujur Rahman",
		companions: []string{"Jony", "Test", "Demo"},
	}
	fmt.Println(aDoctor.companions[1])
}
