//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Print the sequence number and each word of the string
// 1: "Things"
// 2: "don’t"
// ...
//
// @author Alex Versus 2021
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("Things don’t always work out the first time")

	for i, v := range words {
		fmt.Printf("%2d: %q\n", i+1, v)
	}
}
