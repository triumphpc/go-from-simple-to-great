//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 5.2: Arrays
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint
func main() {
	// declaration of array
	var newArray [5]string

	newArray[0] = "a"
	newArray[1] = "b"
	newArray[2] = "c"

	fmt.Println(newArray)

	// null values
	var newArrValues [3]uint8
	fmt.Println(newArrValues) // 0 0 0

	newArrValues[1]++
	newArrValues[1]++
	newArrValues[2]++
	fmt.Println(newArrValues) // 0 2 1

	// arrays with defined  values
	var newInt = [3]int{4, 5, 6}
	var newStr = [3]string{"a", "b", "c"}
	fmt.Println(newInt, newStr) // [4 5 6] [a b c]

	// short declaration
	newShort := [3]int{5, 6, 7}
	fmt.Println(newShort) // [5 6 7]

	newShort2 := [3]int{
		5,
		6,
		7,
	}
	fmt.Println(newShort2) // [5 6 7]

	//dynamic length
	dArr := [...]string{"f", "s", "t", "f"}
	dArr2 := [4]string{"f", "s", "t", "f"}
	fmt.Printf("%T - %[1]v\n", dArr) // [4]string - [f s t f]

	// compilation error
	//fmt.Println(dArr == newStr)

	// normal compare
	fmt.Println(dArr == dArr2) // true

	// key-value declaration
	type DaysOfTheWeek int
	const (
		mon DaysOfTheWeek = iota
		tue
		thu
		fort
		fri
		sat
		sun
	)

	days := [...]string{mon: "one", tue: "two", sat: "six"}
	fmt.Println(days)

	// set null values
	r := [...]int{10: -10}
	fmt.Println(r)

	// error index
	//fmt.Println(r[-5])

	// in for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// use range
	for index, value := range days {
		//fmt.Println(value) // error
		fmt.Println(index, value)
	}

	for _, value := range days {
		//fmt.Println(value) // error
		fmt.Println(value)
	}

	//n-dimensional arrays
	var next [2][5]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			next[i][j] = i + j
		}
	}
	fmt.Println("2d: ", next)

}
