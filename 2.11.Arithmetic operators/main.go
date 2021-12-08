// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.11: Arithmetic operators
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint for program
func main() {
	// Arithmetic operators
	var intExample, intExample2 int = 5, 10
	var floatExample, floatExample2 float64 = 10.9, 21.3

	// Addition
	fmt.Println("Addition")
	fmt.Println(intExample + intExample2)
	fmt.Println(floatExample + floatExample2)

	// Subtraction
	fmt.Println("Subtraction")
	fmt.Println(intExample - intExample2)
	fmt.Println(floatExample - floatExample2)
	// Negation
	fmt.Println(intExample - (-intExample2))
	// Difference
	fmt.Println(intExample - intExample2)

	// Multiplication
	fmt.Println("Multiplication")
	fmt.Println(intExample * intExample2)
	fmt.Println(floatExample * floatExample2)
	// Wrong
	//var res = intExample * floatExample2
	//fmt.Println(res)

	// Quotient
	fmt.Println("Quotient")
	fmt.Println(intExample2 / intExample)
	fmt.Println(-intExample2 / intExample)
	fmt.Println(floatExample / floatExample2)
	// inaccurate values
	fmt.Printf("%.30f\n", 1.0/10)

	// Remainder
	fmt.Println("Remainder")
	fmt.Println(11 % 5)

	// Increment and Decrement
	fmt.Println("Increment and Decrement")
	var number int
	number = number + 5
	fmt.Println(number)
	// Increment it's same
	number += 5
	fmt.Println(number)

	number++
	fmt.Println(number)

	number++
	// Wrong
	//fmt.Println(5 + (number--))
	//fmt.Println(5 + --number)

}
