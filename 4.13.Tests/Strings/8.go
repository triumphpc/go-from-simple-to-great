/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Complete the code to output a not ascii unicode string and code
--
var u rune
u = 'Ø'
--

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	var u rune
	u = 'Ø'
	fmt.Printf("%c (%[1]v)", u)
}
