package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}
type Bird struct {
	Animal
	SpeedKPH float32
	canFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Bangladesh"
	b.SpeedKPH = 150.5
	b.canFly = true
	fmt.Println(b.Name)
}
