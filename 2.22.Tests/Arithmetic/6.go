//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Calculate area of sphere and volume
// https://en.wikipedia.org/wiki/Sphere#Surface_area
// https://en.wikipedia.org/wiki/Sphere#Enclosed_volume
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"math"
)

func main() {
	r := 3.4

	// area of sphere  4 * π * r²
	a := 4 * math.Pi * math.Pow(r, 2)
	fmt.Printf("radius: %g and area: %.2f\n", r, a)

	// volume of sphere
	v := (4 * math.Pi * math.Pow(r, 3)) / 3
	fmt.Printf("radius: %g and volume: %.2f\n", r, v)
}
