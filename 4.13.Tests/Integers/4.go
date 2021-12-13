//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. What the following code will output?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var x uint16 = 65535
	x++

	fmt.Printf("This is type '%T' for value '%[1]v'", x)
}
