/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Its the same values?
2. If not, please fix it

------------
var (
	x uint8
	y uint16
)

x, y = 255, 355
x += 100

fmt.Println(x == y)
--------------


@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	var (
		x uint16
		y uint16
	)

	x, y = 255, 355
	x += 100

	fmt.Println(x == y)
}
