//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. We have loop:
// for i := 1; i < 10; i = i + 2 {}
// 2. Output sum value of i in every iteration
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var s = 0
	for i := 1; i < 10; i = i + 2 {
		s += i
	}

	fmt.Println("Sum:", s)
}
