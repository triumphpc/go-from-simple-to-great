// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.19: Visibility blocks
//
// @author Alex Versus 2021
// www.alexversus.com

// package block
package main

import "fmt"

// function block
func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, у) // Compiling error

}

func f() int {
	return 1
}

func g(x int) int {
	return x

}
