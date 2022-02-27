package main

import "fmt"

type VerTex struct {
	Latitude, Longitude float64
}

var m = map[string]VerTex{
	"JavaTpoint":     VerTex{40.00546, -54.564641},
	"TurorialsPoint": VerTex{30.00546, -34.564641},
}

func main() {
	fmt.Println(m)

}
