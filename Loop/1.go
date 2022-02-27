//For Loop Nested
package main

import "fmt"

func main() {
	var a int
	var b int

	for a = 0; a < 3; a++ {
		for b = 3; b > 0; b-- {
			fmt.Print(a, "  ", b, "\n")
		}

	}
}
