/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. What output of this code?
-----------------
var x float64 = 0

for i := 0; i < 11 ; i++ {
	x += 0.10
}

fmt.Println(x)
-----------------

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	var x float64 = 0

	for i := 0; i < 11; i++ {
		x += 0.10
	}

	fmt.Println(x)
}
