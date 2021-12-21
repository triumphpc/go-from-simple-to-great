//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function to warn that text contains a vowel
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"strings"
)

func main() {
	run()
}

// Run check vowel
func run() {
	fmt.Println("Enter text:")

	var text string
	fmt.Scanln(&text)

	// Check has a vowel
	for _, v := range strings.ToLower(text) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			fmt.Println("Contains a vowel.")
			return
		}
	}
	fmt.Println("Not contain a vowel.")
}
