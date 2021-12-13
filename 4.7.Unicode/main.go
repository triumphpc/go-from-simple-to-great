//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.7: Unicode
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"unicode/utf8"
)

// Entrypoint
func main() {
	s := "Hello, 世界"
	fmt.Println(s)

	fmt.Println(len(s))                    // len string in bytes
	fmt.Println(utf8.RuneCountInString(s)) // count of runes

	// decode
	for i := 0; i < len(s); {
		r, sz := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%q\t%d\n", i, r, r) // index rune unicode

		i += sz
	}

	// incorrect rune
	fmt.Println(string(123456789)) // ?
}
