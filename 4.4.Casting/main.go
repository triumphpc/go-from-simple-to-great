//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Example 4.4: Casting
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
	// conversion example
	var a int32 = 5
	var b int64 = 10
	//var c int = a + b // error
	var c int = int(a) + int(b) // correct
	fmt.Println(c)

	// output type
	fmt.Printf("data type: %T\n", a)
	fmt.Printf("data type: %T\n", b)

	// transfer
	var n int64 = 129
	var l = int8(n)
	fmt.Println(l) // -127

	f := 3.141 // float64
	i := int(f)
	fmt.Println(f, i)   // 3.141 3
	fmt.Println(int(f)) // 3

	// float convert
	var z int64 = 34
	var z2 float64 = float64(z)
	fmt.Printf("%.2f\n", z2) // 34.00

	var f64 float64 = 11.9
	var iex int = int(f64)
	fmt.Printf("f64 = %.2f\n", f64)
	fmt.Printf("iex = %d\n", iex)

	// round to ceil
	f64ToInt := math.Round(f64)
	fmt.Println(f64ToInt)

	// division
	d := 10
	del := 3
	fmt.Println("divide conversion", d/del)        // 3
	fmt.Println("divide conversion float", 10.0/3) // 3.33

	// output format
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // 438 666 0666
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x\n", x) // 3735928559 deadbeef 0xdeadbeef

	ascii := 'а'
	Unicode := 9733
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 1072 а 'а'
	fmt.Printf("%d %[1]c %[1]q\n", Unicode) // 9733 ★ '★'
	fmt.Printf("%d %[1]q\n", newline)       // 10 '\n'

}
