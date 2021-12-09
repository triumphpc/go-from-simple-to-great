//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write a program that takes two arguments as input
// 2. The first argument is a command, the second argument is string
// 3. Command - can be any command over the string
// 4. Output result
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Please enter <command> <string>")
		return
	}

	switch args[1] {
	case "up":
		fmt.Println(strings.ToUpper(args[2]))
	case "down":
		fmt.Println(strings.ToLower(args[2]))
	case "trim":
		fmt.Println(strings.Trim(args[2], " "))
	default:
		fmt.Println("Unknown cmd")
	}
}
