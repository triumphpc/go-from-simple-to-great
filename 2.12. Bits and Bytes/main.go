// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.12: Bits
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint for program
func main() {
	// 2^1 = 2
	fmt.Printf("%02b\n", 0)
	fmt.Printf("%02b\n", 1)
	// 2^2 = 2
	fmt.Printf("%02b\n", 2)
	fmt.Printf("%02b\n", 3)

	// bytes
	var byteSize byte

	byteSize = 2
	fmt.Printf("%08b = %d\n", byteSize, byteSize)

	byteSize = 255
	fmt.Printf("%08b = %d\n", byteSize, byteSize)

}
