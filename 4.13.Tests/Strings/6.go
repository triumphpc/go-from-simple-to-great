/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Remove extra spaces from string variable in this code

@author Alex Versus 2021
www.alexversus.com
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	x := `


text here
and text here
and here



`

	fmt.Printf("This is type '%T' for value '%[1]v'", strings.TrimSpace(x))
}
