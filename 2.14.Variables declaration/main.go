// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.14: Variable declaration
//
// @author Alex Versus 2022
// www.alexversus.com

package main

import "fmt"

var (
	packInt      int    = 32
	packStr      string = "String var"
	packageLevel int    = 55
)

// short in package
//packShortDeclaration := 123

// Entrypoint for program
func main() {
	// variable declaration
	var myAge int8
	var myIntroduction string

	myAge = 35
	myIntroduction = "Super"
	fmt.Println(myAge, myIntroduction)

	// together in zero
	var v, m int8
	fmt.Println(v, m)

	// together and set
	var v2, m2 int8 = 5, 10
	// Using variables
	fmt.Println(v2, m2)

	//Reassign
	v2 = 10
	fmt.Println(v2) // 10

	// Null-value
	var v3 int64
	var sv string
	var bv bool
	fmt.Println(v3, sv, bv) // 0 _ false

	// Short variable declaration
	quik := 85
	quikS := "This in shadowing"
	fmt.Printf("%T : %v\n", quik, quik) // 85
	fmt.Printf("%T : %v\n", quikS, quikS)
	// redeclaration
	bv, newBv := false, 123
	fmt.Printf("%T : %v\n", newBv, bv)

	// classic
	var varVal int8
	varVal = 123
	fmt.Printf("%T : %v\n", varVal, varVal)

	// one line
	var oneVarVal int8 = 5
	fmt.Printf("%T : %v\n", oneVarVal, oneVarVal)

	// out package var
	fmt.Println(packInt, packStr)

	// redeclare
	var packageLevel = 66
	packageLevel = 77
	fmt.Println(packageLevel)

	// multiple variables
	var first, second int = 4, 7
	fmt.Println(first, second)

	//blank identifier
	var forBlank int = 4
	_ = forBlank

	// assign to variables
	forBlank, first = packageLevel, quik

	// from function
	var fromFn = someFunc()
	fmt.Println(fromFn)

	// swapping
	var from, to int = 5, 10
	from, to = to, from
	fmt.Println(from, to) // 10, 5
}

/**
Function for assigment
*/
func someFunc() string {
	return "From function"
}
