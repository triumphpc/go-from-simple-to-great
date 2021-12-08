//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Calculate area of circle
// https://en.wikipedia.org/wiki/Area_of_a_circle
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
	var (
		r = 17.
		a float64
	)

	// S=πr2
	a = math.Pi * math.Pow(r, 2)
	fmt.Printf("radius: %g area: %.2f\n", r, a)

}
