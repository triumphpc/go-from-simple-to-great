//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. What is the output of this program?
// x, y := 5, 1.5
// x++
// x++
// x--
// y--
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	x, y := 5, 1.5
	x++
	x++
	x--
	y--

	fmt.Println(float64(x) * y) // 3

}
