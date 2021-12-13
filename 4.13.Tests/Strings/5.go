/**
Go (Golang) From simple to great. The Complete Developer's Guide.

1. Write a program that upper case string from args


Example: go run 5.go go is "good" language
Output: GO IS GOOD LANGUAGE

@author Alex Versus 2021
www.alexversus.com
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	lenArg := len(args[1:])
	out := ""

	for i := 1; i <= lenArg; i++ {
		out += strings.ToUpper(args[i]) + " "
	}

	fmt.Println(out)
}
