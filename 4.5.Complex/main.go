//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.5: Complex
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"math"
)

// Entrypoint
func main() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i

	x2 := 1 + 2i
	y2 := 3 + 4i

	fmt.Println(x + y)   // 4 + 6i
	fmt.Println(x2 + y2) // 4 + 6i

	fmt.Println(x * y) //(5 + 10i)

	// real part
	fmt.Println(real(x * y)) // 5
	// complex part
	fmt.Println(imag(x * y)) // 10

	// cast
	var a = complex(1, -2) // complex128
	fmt.Println(a)
	xb := float32(math.Cos(math.Pi / 2)) // float32
	fmt.Println(xb)
	var c64 = complex(5, -xb) // complex64
	fmt.Println(c64)
	//var s uint = complex(1, 0) // untyped complex 1 + 0i can be converted to uint
	//_ = complex(1, 2<<s) // invalid: 2 assumes floating point, cannot shift
	var fl = real(c64) // float32
	fmt.Println(fl)
	var im = imag(a) // float64
	fmt.Println(im)
	//_ = imag(3 << s) // invalid: 3 assumes complex, cannot shift

}
