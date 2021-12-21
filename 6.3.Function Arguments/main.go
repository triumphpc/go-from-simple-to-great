//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.3: Functions params and arguments
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
)

// Entrypoint
func main() {
	// as arguments
	fmt.Println(example(1, 2, 3))

	// with current slice
	var s = []int{1, 2, 3, 4, 5}
	fmt.Println(example(s...))
}

func example(params ...int) int {
	t := 0
	for v := range params {
		t += v
	}
	return t
}
