//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Detect capital letter or not
// 2. Print this for each letter of the string
//
// ------------------------
// s := "this is String with LETTERS"
//
// ------------------------
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "this is String with LETTERS"

	for _, v := range s {
		if unicode.IsUpper(v) {
			fmt.Printf("%c is capital\n", v)
		} else {
			fmt.Printf("%c is lower\n", v)
		}
	}
}
