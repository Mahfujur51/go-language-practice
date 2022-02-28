//Anonymous Structure

package main

import "fmt"

func main() {
	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Mahfujur",
		branch:    "Mirpur",
		language:  "Go",
		Particles: 60,
	}
	fmt.Println(Element)
}
