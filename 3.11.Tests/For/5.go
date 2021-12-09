//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Print the sequence number and each character of the string
// 1: 'T'
// 2: 'h;
// ...
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
)

func main() {
	words := "Things don’t always work out the first time"

	for i, v := range words {
		fmt.Printf("%2d: %q\n", i+1, v)
	}
}
