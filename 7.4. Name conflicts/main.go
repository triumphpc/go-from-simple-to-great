//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.4: Name conflict resolving
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Defined type
type MyStruct struct{ x, y int }

type AnotherType MyStruct

// Entrypoint
func main() {
	s1 := MyStruct{1, 2}
	s2 := MyStruct{5, 6}

	// call function
	fmt.Println(ExampleName(s1, s2)) // 48

	// call method
	fmt.Println(s1.ExampleName(s2)) // 48

	a1 := AnotherType{2, 3}
	a2 := AnotherType{7, 8}
	fmt.Println(a1.ExampleName(a2)) // 25

}

// Traditional function name
func ExampleName(x, y MyStruct) int {
	return (x.x + y.x) * (x.y + y.y)
}

// Method  with the same name and implementation
func (m MyStruct) ExampleName(y MyStruct) int {
	return (m.x + y.x) * (m.y + y.y)
}

// Method  with the same name for other type
func (m AnotherType) ExampleName(y AnotherType) int {
	return (m.x - y.x) * (m.y - y.y)
}
