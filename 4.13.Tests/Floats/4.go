//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Output a float64 value with a width of 5 and precision 3, for example "00123.457"
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	x := 123.456789

	fmt.Printf("This is type '%T' for value '%09.3[1]f'", x)
}
