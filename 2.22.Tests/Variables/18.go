//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Get variables from Args
// 2. Print count of args
// 3. Print first arg
// 4. go run 18.go hello
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Count of args: ", len(args)-1)
	fmt.Println("Path: ", args[0])
	fmt.Println("First arg: ", args[1])
}
