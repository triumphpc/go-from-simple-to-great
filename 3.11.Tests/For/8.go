//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write a program that will search for the occurrence of a word in string
// 2. If a word is found - displays "Word found", otherwise "Word not found"
// 3. Be sure to use labels in the program. Where exactly is up to you,
// but the program should not do unnecessary actions when finding a word
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s = "Beauty and wisdom are seldom found together."
	words := strings.Fields(s)
	input := os.Args[1:]

	var f = false
out:

	for _, q := range input {
		for _, w := range words {
			if q == w {
				fmt.Println("Word found")
				f = true
				break out
			}
		}
	}
	if f == false {
		fmt.Println("Word not found")
	}
}
