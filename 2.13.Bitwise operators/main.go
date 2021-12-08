// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.13: Bitwise operators
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
)

const (
	UPPER = 1 // upper case
	LOWER = 2 // lower case
	CAP   = 4 // capitalizes
	REV   = 8 // reverses
)

// Entrypoint for program
func main() {
	// AND
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Result AND %d\n", c)

	// clear the last 4 least significant bits
	var x uint8 = 172 // 172 = 10101100
	var y uint8 = 240 // 240 = 11110000
	x = x & y         // x   = 10100000

	fmt.Printf("Result AND %d\n", x)
	// short variant
	x &= y // x   = 10100000
	fmt.Printf("Result two again %d\n", x)

	// example check odd/even
	var num = 123
	if num&1 == 1 {
		fmt.Printf("%d is odd\n", num)
	} else {
		fmt.Printf("%d is even\n", num)
	}

	num = 120
	if num&1 == 1 {
		fmt.Printf("%d is odd\n", num)
	} else {
		fmt.Printf("%d is even\n", num)
	}

	// OR
	var f uint8 = 0
	fmt.Printf("OR: %d = %08b\n", f, f)
	f |= 197
	fmt.Printf("OR: %d = %08b\n", f, f)
	f |= 196
	fmt.Printf("Result OR: %d = %08b\n", f, f)

	// configuration example
	fmt.Println("\n----")
	checkConfig(UPPER | CAP)
	fmt.Println("----")
	checkConfig(REV | CAP)

	//XOR
	a = 60    /* 60 = 0011 1100 */
	b = 13    /* 13 = 0000 1101 */
	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Result XOR: %d = %08b\n", c, c)

	//comparison of signs
	k, l := -74, 35
	fmt.Println("It's the same sign'?", (k^l) >= 0)
	k, l = 74, 35
	fmt.Println("It's the same sign'?", (k^l) >= 0)

	// NOT XOR
	var by byte = 15
	fmt.Println("Result NOT XOR")
	fmt.Printf("%08b\n", by)
	fmt.Printf("%08b\n", ^by)

	//AND_NOT
	fmt.Println("Result AND_NOT")
	var an byte = 25
	var anb byte = 15
	fmt.Printf("%08b\n", an)
	fmt.Printf("%08b\n", anb)
	an &^= anb
	fmt.Printf("%08b\n", an)

	// << >>
	var arr int8 = 5
	fmt.Printf("%08b\n", arr)
	fmt.Printf("%08b\n", arr<<1)
	fmt.Printf("%08b\n", arr<<2)
	fmt.Printf("%08b\n", arr<<3)

	a = 70
	fmt.Println("Result <<>>")
	// division by 2
	fmt.Printf("%d\n", a>>1)
	//  multiply the value by 4
	fmt.Printf("%d\n", a<<2)

	// set fifth bit to 1
	ab := 9
	fmt.Printf("%08b\n", ab)
	ab = ab | (1 << 4)
	fmt.Printf("%08b\n", ab)

	// check this
	if ab&(1<<4) != 0 {
		fmt.Println("yup")
	}

	// reset bit again
	ab = ab &^ (1 << 4)
	fmt.Printf("%08b\n", ab)
	if ab&^(1<<4) != 0 {
		fmt.Println("ooh")
	}

}

func checkConfig(setParam byte) {
	// query config bits
	if (setParam & UPPER) != 0 {
		fmt.Println("It's upper logic")
	}
	if (setParam & LOWER) != 0 {
		fmt.Println("It's lower logic")
	}
	if (setParam & CAP) != 0 {
		fmt.Println("It's cap logic")
	}
	if (setParam & REV) != 0 {
		fmt.Println("It's rev logic")
	}
}
