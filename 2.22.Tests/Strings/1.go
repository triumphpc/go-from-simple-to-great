//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. What's the output?
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	s := "x"
	st := 'x' + 'A' - 'a'
	stNew := s[0] + 'A' - 'a'

	fmt.Printf("%c\n", st)
	fmt.Printf("%c\n", stNew)
}
