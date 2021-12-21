//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.12: Lifetime of variables
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
)

var gl *int

// Entrypoint
func main() {
	for i := 0; i < 10; i++ {
		x := i * 2
		y := x * 2
		fmt.Println(y)
	}

	// example
	f()
	fmt.Println(gl)
	fmt.Println(*gl)

	// second example
	g()
}

func f() {
	var x int
	x = 1
	gl = &x
}

func g() {
	y := new(int)
	*y = 1
}
