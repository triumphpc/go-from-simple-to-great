/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Rewrite your program so that a string variable is defined as an unsigned plus-sign string literal (use raw string).
2. The value of the string variable must contain several new lines
--------
x := "text here" +
	" and text here"

fmt.Printf("This is type '%T' for value '%[1]v'", x)
--------



@author Alex Versus 2021
www.alexversus.com
*/

package main

import "fmt"

func main() {
	x := `text here
and text here
and here
`

	fmt.Printf("This is type '%T' for value '%[1]v'", x)
}
