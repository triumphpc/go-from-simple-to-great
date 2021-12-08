//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Declare and print a variable with all string, int, bool, float types
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	// integers
	var in int = 123
	var in8 int8
	var in16 int16
	var in32 int32
	var in64 int64

	fmt.Printf("var %v with type  %T\n", in, in)
	fmt.Printf("var %v with type  %T\n", in8, in8)
	fmt.Printf("var %v with type  %T\n", in16, in16)
	fmt.Printf("var %v with type  %T\n", in32, in32)
	fmt.Printf("var %v with type  %T\n", in64, in64)

	// floats
	var fl32 float32 = 3.14
	var fl64 float64
	fmt.Printf("var %v with type  %T\n", fl32, fl32)
	fmt.Printf("var %v with type  %T\n", fl64, fl64)

	// bool
	var bl bool
	fmt.Printf("var %v with type  %T\n", bl, bl)

	// string types
	var st string = "Yup"
	var rn rune // also a numeric type
	var bt byte // also a numeric type
	fmt.Printf("var %v with type  %T\n", st, st)
	fmt.Printf("var %v with type  %T\n", rn, rn)
	fmt.Printf("var %v with type  %T\n", bt, bt)

	str := "iphone"
	fmt.Printf("var %v with type  %T\n", str, str)
}
