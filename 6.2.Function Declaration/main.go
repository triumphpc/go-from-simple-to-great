//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.2: Functions declaration
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"strconv"
	"time"
)

// declaration unnamed args func in struct
type export struct {
	f      func(time.Time, time.Time, string)
	folder string
}

// Entrypoint function declaration
func main() {
	fmt.Println(example(1, 2))
	fmt.Println(unnamedArgs(1, 2))
}

func example(x, y int) (st string) {
	x += y
	return "return string" + strconv.Itoa(x)
}

// Second example
func exampleTwo() string {
	return "return string"
}

// Example unnamed args function
func unnamedArgs(int, int) int {
	return 1
}
