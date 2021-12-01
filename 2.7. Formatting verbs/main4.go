// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.7: Formatting verbs
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint for program
func main() {
	// string
	fmt.Printf("It's '%s' string and '%d' digit\n", "string", 123)
	// float
	fmt.Printf("It's '%f' float\n", 123.456)
	// decimal integer
	fmt.Printf("It's '%d' decimal\n", 12345)
	// boolean
	fmt.Printf("It's '%t' boolean\n", false)
	// arbitrary value
	fmt.Printf("It's '%v %v %v %v' arbitrary values\n", 123, true, "\t", 1.2)
	// arbitrary value in go
	fmt.Printf("It's '%#v %#v %#v %#v' arbitrary values in GO\n", 123, true, "\t", 1.2)
	// type of value
	fmt.Printf("It's '%T %T %T %T' types of values\n", 123, true, "\t", 1.2)
	// percent
	fmt.Printf("It's '%%' percent\n")
	// quote
	fmt.Printf("It's %q quote\n", "Go (Golang) From simple to great")
	// warning
	fmt.Printf("It's %d quote\n", 123.45)

	var (
		country = "Russia"
		area    = 17130000
	)
	// Use index for variables
	fmt.Printf(
		"Area of %v is %v kms squared. Think! %[2]v kms squared! %[1]v Amazing.\n",
		country, area,
	)

	// Exsample with Escape sequences
	fmt.Printf("It's escape \t \" \n \\ sequences exsmples\n")
}
