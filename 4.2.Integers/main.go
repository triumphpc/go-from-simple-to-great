//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.2: Integers
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"time"
)

// Entrypoint
func main() {
	// synonyms
	var r rune = 'R'
	var i32 int32 = 'R'

	if r == i32 {
		fmt.Println("It's synonyms")
	}

	var b byte = 1
	var ui8 uint8 = 1
	if b == ui8 {
		fmt.Println("It's synonyms")
	}

	// integers
	var ii uintptr = 5
	fmt.Println(ii)

	//remainder of the division
	fmt.Println(-5 % 3)  // -2
	fmt.Println(-5 % -3) // -2

	fmt.Println(5 / 4)     // 1
	fmt.Println(5.0 / 4.0) // 1.25

	//overflow
	var u8 uint8 = 255
	fmt.Println(u8, u8+1, u8*u8) // 255 0 1

	var i8 int8 = 127
	fmt.Println(i8, i8+1, i8*i8) // 127 -128 1

	f := time.Unix(12622780800, 0)
	fmt.Println(f)

}
