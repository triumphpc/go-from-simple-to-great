//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 8.6: Debugging
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"runtime"
)

// Entrypoint
func main() {
	var s int = 1

	fmt.Println(s)

	s++
	// run breakpoint
	runtime.Breakpoint()

	someFunc(&s)

	fmt.Println(s)
}

func someFunc(i *int) {
	*i = 5
}
