//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Convert variables to get output 20.7
// INPUT:
// x, y := 13, 7.7
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	// you have input
	x, y := 13, 7.7

	//first we convert y to int
	z := int(y) // int 7
	t := z + x  // int 20

	fmt.Print("Output: ", float64(t)+y-float64(z)) // 20.7
}
