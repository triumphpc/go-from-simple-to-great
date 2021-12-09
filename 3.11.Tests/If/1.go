//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Implement input and read from argv
// 2. If the first parameter more than 50 - print: You are old for such games
// 3. If the first parameter less than 18 - print: Permission denied
// 4. Else print - Welcome!
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var args = os.Args

	if len(args) > 1 {
		if i, err := strconv.Atoi(args[1]); err == nil {
			if i > 50 {
				fmt.Println("You are old for such games")
			} else if i < 18 {
				fmt.Println("Permission denied")
			} else {
				fmt.Println("Welcome")
			}
		}
	}
}
