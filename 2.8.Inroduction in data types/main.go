//  Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.8: Data types
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// Entrypoint for program
func main() {
	// string
	var stringExample = "It's string example"

	// Rune output example
	var runeExample = 'A'
	var runeNewLine = '\n'
	var runeTab = '\t'

	// Boolean example
	var boolExample = false

	// Boolean example
	//var boolTrueExample = true

	fmt.Printf("%T : %v\n", stringExample, stringExample)
	fmt.Printf("%T : %v\n", runeExample, runeExample)
	fmt.Printf("%T : %v\n", runeNewLine, runeNewLine)
	fmt.Printf("%T : %v\n", runeTab, runeTab)
	fmt.Printf("%T : %v\n", boolExample, boolExample)

	// Numbers
	var kms int = 100
	var weight float64 = 970.5

	fmt.Printf("To this city %v km : %T\n", kms, kms)
	fmt.Printf("This car weight %v kg : %T\n", weight, weight)

	// Using
	var someTextHere string = "Amazing language"

	// Wrong using of types
	//fmt.Println(strings.TrimSpace(weight))
	fmt.Println(strings.TrimSpace(someTextHere))

	// Wrong using of types
	//fmt.Println(strings.TrimSpace(weight))
	fmt.Println(math.Ceil(weight))

	// get types from reflection
	fmt.Println(reflect.TypeOf(weight))
	fmt.Println(reflect.TypeOf(someTextHere))
	fmt.Println(reflect.TypeOf(boolExample))
	fmt.Println(reflect.TypeOf(runeExample))
	fmt.Println(reflect.TypeOf(stringExample))

}
