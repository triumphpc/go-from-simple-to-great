//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.9: Constants
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"time"
)

const packageConst = "Package const value"
const DayInWeek = 7

const exampleScope = "package scope"

// Entrypoint
func main() {
	const c string = "It's constant value"
	fmt.Println(c)

	// error
	//c = "New value"

	// unnamed constants
	fmt.Println("It's unnamed constant", 555)

	// declaration
	//const fromFunc float64 = math.Round(100) // runtime declaration
	//var ex = 10
	//const fromVar int ex // runtime declaration

	const fromBuildIn int = len("from string")
	fmt.Println(fromBuildIn)

	// declaration from expresion
	const fromExpression int = 5 * 10
	fmt.Println(fromExpression)

	// using const from package level
	fmt.Println(packageConst)

	// local const with type
	const hasNotDelay time.Duration = 0
	const timeout = 10 * time.Second

	fmt.Printf("%T %[1]v\n", hasNotDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Second)

	// compile time error
	const first int = 5
	const second int = 0
	//fmt.Println(first / second) // division by zero

	// multiple declaration
	const constOne, constTwo int = 1, 2

	// block declaration
	const (
		a = 1
		b
		v = 3
	)

	fmt.Println(a, b, v)

	// block with expression
	const (
		m = 1 + 1
		n
	)
	fmt.Println(m, n)

	// example iota
	const (
		oneDay int = iota
		twoDay
		threeDay
	)

	fmt.Println(oneDay, twoDay, threeDay)

	// uniq binary flags
	const (
		FlagOne = 1 << iota
		FlagTwo
		FlagThree
	)

	// blank identifier in iota
	//const (
	//	targetOne = 5
	//	targetTwo = 8
	//	targetThree = 10
	//	targetFour = 11
	//)

	// blank identifier in iota
	const (
		targetOne int = (5 + iota)
		_
		_
		targetTwo
		_
		targetThree
		targetFour
	)

	fmt.Println(targetOne, targetTwo, targetThree, targetFour)

	// untyped constant
	const Pi = 3.14

	// use different types
	const cDeclaration = 5 * 3.25
	fmt.Printf("%T - %[1]v", cDeclaration)

	// untyped unless constants
	const piEx = 3.14 * 5
	fmt.Println(piEx)

	// same type
	const intC int = 5
	var intV = 5
	fmt.Println(intC * intV) // only the same type

	// with different types
	const min = 66
	var i1 int = min
	var i2 float64 = min
	var i3 byte = min
	var i4 rune = min
	fmt.Println(i1, i2, i3, i4)

	// untyped const in expression
	uEx := time.Second * 15
	fmt.Printf("%T - %[1]v\n", uEx)
	//from var
	iExInt := 15
	//fmt.Println(time.Second * iExInt)  // error
	fmt.Printf("%T - %[1]v\n", time.Second*time.Duration(iExInt)) // int64 = time.Duration

	const cUx = 15
	fmt.Printf("%T - %[1]v\n", time.Second*cUx) // int64 = time.Duration

	// shadowing scope const
	const packageConst = "package scope new value"
	fmt.Println(packageConst)
	getPackageScopeConst()

}

func getDays(weekCount int) int {
	return weekCount * DayInWeek
}

func getPackageScopeConst() {
	fmt.Println(packageConst)

}
