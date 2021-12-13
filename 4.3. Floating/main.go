//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.3: Floating
//
//
// @author Alex Versus 2021
//

package main

import (
	"fmt"
	"math"
)

// Entrypoint
func main() {
	d := 365.2425
	fmt.Println(d)
	var d2 = 365.2425
	fmt.Println(d2)
	var d3 float64 = 365.2425
	fmt.Println(d3)

	// floating precision
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64) // 3.141592653589793
	fmt.Println(pi32) // 3.1415927

	//var f1 = .707
	//var f2 = 1.

	// scientific number format
	var a = 4.422434129e23
	var b = 5.6606957e34
	fmt.Println(a)
	fmt.Println(b)

	// output
	for x := 0; x < 8; x++ {
		fmt.Printf("х = %d, ех = %8.3f, g=%g, e=%e\n", x, math.Exp(float64(x)), float64(x), float64(x))
	}

	// nil-value
	var p float64
	p2 := 0.0
	fmt.Println(p, p2) // 0

	// width
	fmt.Printf("%05.2f\n", .55) // 00.55

	// rounding accuracy
	t := 1.0 / 3.0
	fmt.Println(t + t + t) // 1

	m := 0.2
	m += 0.1
	fmt.Println(m) // 0.30000000000000004

	c := 21.0
	// division before
	f := (c / 5.0 * 9.0) + 32
	// multiplication before
	f2 := (c * 9.0 / 5.0) + 32

	fmt.Println(f, "° F")  // 69.80000000000001 ° F
	fmt.Println(f2, "° F") // 69.8 ° F

	// comparison
	pp := 0.1
	pp += 0.2
	fmt.Println(pp == 0.3) // false

}
