/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Complete the code to output a unicode string
--
var r rune = 'A'
fmt.Println("This is string in unicode:", r)
--

@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	var r rune = 'A'
	fmt.Println("This is string in unicode:", string(r))
}
