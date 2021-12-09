//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write a program that accepts any number of arguments as input
// 2. Arguments can be any characters
// 3. The program must remove all spaces and numbers from the input arguments and convert to one single string
// 4. Print the result
// Example:
// go run 14.go fs FD 34 *fff 43F
// fsFD*fffF
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var out = ""
	for _, arg := range os.Args[1:] {
		for _, a := range arg {
			if unicode.IsDigit(a) {
				// skip numbers
				continue
			}
			out += string(a)
		}
	}
	// delete all whitespaces
	fmt.Println(strings.ReplaceAll(out, " ", ""))
	fmt.Println()

}
