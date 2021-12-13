//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.10: Defined types
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type USD float64
type EUR float64 // ratio 1.3

// Entrypoint
func main() {
	// new types with base type int32
	type Week int32
	type Month int32

	// convert int32 type to Week
	var semesterPeriod = Month(6)
	var examPeriod = Week(2)

	fmt.Printf("%T - %[1]v\n", semesterPeriod)
	fmt.Printf("%T - %[1]v\n", examPeriod)

	// it's work, but not correct - it's the same
	fmt.Printf("%T - %[1]v\n", Month(Week(24))) // main.Month - 24 it's not 6 months
	fmt.Printf("%T - %[1]v\n", Week(Month(6)))  //main.Week - 6 it's not 24 weeks

	// convert base type to needed value
	var semesterPeriodInWeeks = Week(Month(6) * 4)
	var examPeriodInMonths = Month(Week(2) / 4)

	fmt.Printf("%T - %[1]v\n", semesterPeriodInWeeks) // 24
	fmt.Printf("%T - %[1]v\n", examPeriodInMonths)    // 0 because it's int

	// arithmetic operators
	fmt.Printf("%T - %[1]v\n", Week(5)-Week(2)) // main.Week - 3
	fmt.Printf("%T - %[1]v\n", Week(5)+Week(2)) // main.Week - 7
	fmt.Printf("%T - %[1]v\n", Week(5)*Week(2)) // main.Week - 10

	type description string
	fmt.Printf("%T - %[1]v\n", description("some text") == "some text") //bool - true
	fmt.Printf("%T - %[1]v\n", description("some text") > "yanki")      //bool - false
	fmt.Printf("%T - %[1]v\n", description("some text")+" yanki")       //main.description - some text yanki

	// literal together
	fmt.Printf("%T - %[1]v\n", Week(5)+2) // main.Week - 7
	fmt.Printf("%T - %[1]v\n", Week(5)*2) // main.Week - 10

	// same base type
	//fmt.Printf("%T - %[1]v\n", Week(5) * Month(1)) // error

	currentBalance := USD(20)

	// add in EUR
	currentBalance += USD(EUR(10) * 1.3)
	currentBalance += toUSD(10)

	fmt.Printf("%T - %[1]v\n", currentBalance) // main.USD - 46

	// create new defined from defined type
	type newUSD USD // it own defined type with underlying type float 64

}

func toUSD(v EUR) USD {
	return USD(v * 1.3)
}
