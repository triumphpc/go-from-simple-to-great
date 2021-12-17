//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 5.3: Slices
//
//@author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
)

// Entrypoint
func main() {
	// declaration
	colors := []string{"red", "blue", "green", "yellow"}
	fmt.Printf("%T %[1]v\n", colors) // []string [red blue green yellow]

	// use make
	workDays := make([]string, 5)
	fmt.Printf("%T %[1]q\n", workDays) // []string ["" "" "" "" ""]

	// use make with volume
	weekDays := make([]string, 5, 7)
	fmt.Printf("%T %[1]q\n", weekDays) // []string ["" "" "" "" ""]

	// get slice
	fmt.Printf("%T %[1]q\n", colors[1:2]) // []string ["blue"]

	// get slice
	fmt.Printf("%T %[1]q\n", colors[1:]) // ["blue" "green" "yellow"]

	// get full slice
	fmt.Printf("%T %[1]q\n", colors[:]) // []string ["red" "blue" "green" "yellow"]

	// several slices
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	startWeek := days[:3]
	endWeek := days[2:5]

	// change base array
	days[2] = "changed day"
	fmt.Println(startWeek, endWeek) // [monday tuesday changed day] [changed day thursday friday]

	// you can compare with nil
	if startWeek == nil {
		// do something
	}
	// nil examples all equal
	var st []int
	fmt.Printf("%T %[1]q %v %v\n", st, len(st), st == nil)
	st = nil
	fmt.Printf("%T %[1]q %v %v\n", st, len(st), st == nil)
	st = []int(nil)
	fmt.Printf("%T %[1]q %v %v\n", st, len(st), st == nil)
	st = []int{}
	fmt.Printf("%T %[1]q %v %v\n", st, len(st), st == nil) // false

	//change slice
	forChange := []string{"one", "two", "three"}
	fmt.Println(forChange, len(forChange)) // [one two three] 3
	forChange = append(forChange, "four", "five")
	fmt.Println(forChange, len(forChange)) // [one two three four five] 5

	// copy slices example
	var xSt, ySt []string
	for _, value := range forChange {
		ySt = appendString(xSt, value)
		fmt.Printf("%q c=%d\t%v\n", value, len(ySt), ySt)
		xSt = ySt
	}

	// nil for append
	var stapp []int
	stapp = append(stapp, 12)
	fmt.Println(stapp)

}

func appendString(x []string, y string) []string {
	var z []string
	addLen := len(x) + len(y)

	if addLen <= cap(x) {
		// increase slice
		z = x[:addLen]
	} else {
		//create new array
		addCap := addLen
		if addCap < 2*len(x) {
			addCap = 2 * len(x)
		}
		z = make([]string, addLen, addCap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}
