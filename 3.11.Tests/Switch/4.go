//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write program which displays part of the day according to the current time
// Example : go run 4.go
// Good morning!
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"time"
)

func main() {
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
