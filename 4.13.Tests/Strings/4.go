/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Write a program that will strip out all non-rune literal characters and print the result


Example: go run 4.go go is "good" language
Output: goisgoodlanguage

@author Alex Versus 2021
www.alexversus.com
*/

package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	args := os.Args
	lenArg := len(args[1:])
	out := ""

	for i := 1; i <= lenArg; i++ {
		for j := 0; j < len(args[i]); j++ {
			r := rune(args[i][j])
			if utf8.ValidRune(r) {
				out += string(r)
			}
		}
	}

	fmt.Println(out)
}
