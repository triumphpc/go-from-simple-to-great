//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. How output decimal number in bits (with base two)?
//
// @author Alex Versus 2021
// www.alexversus.com
//
package main

import "fmt"

func main() {
	var x uint16 = 10
	x++

	fmt.Printf("This is type '%T' for value '%[1]v' in bits '%[1]b'", x)
}
