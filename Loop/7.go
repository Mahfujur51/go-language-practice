// Go Continue Statement with Inner Loop example:

package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 1
	var b int = 1
	/* do loop execution */
	for a = 1; a < 3; a++ {
		for b = 1; b < 3; b++ {
			if a == 2 && b == 2 {
				/* skip the iteration */
				continue
			}
			fmt.Printf("value of a and b is %d %d\n", a, b)
		}
		fmt.Printf("value of a and b is %d %d\n", a, b)
	}
}
