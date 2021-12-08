/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Convert int type to float using only arithmetic operator
2. Convert int type to float using conversion operator

@author Alex Versus 2021
*/

package main

import "fmt"

func main() {
	r := 5 + 0.1
	fmt.Printf("%v : %T\n", r, r)

	i := 5
	f := 5.5
	r = float64(i) + f
	fmt.Printf("%v : %T\n", r, r)

}
