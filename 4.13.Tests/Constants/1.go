//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Declare a constant with a bit value in a byte
// 2. Declare a constant with a byte value in kilobyte (1000)
// 3. Write an expression that will output the number of bits in 10 kilobytes
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

const (
	bitInByte   uint8  = 8
	byteInKbyte uint16 = 1000
)

func main() {
	fmt.Println(10 * uint32(byteInKbyte) * uint32(bitInByte))
}
