//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Print variable as appears in GO
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var s string
	fmt.Printf("s (%T): %q\n", s, s)
	fmt.Printf("s (%T): %v\n", s, s)
}
