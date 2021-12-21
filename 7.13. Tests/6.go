//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Will the code below work?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Defined type
type MyStruct struct{ x, y int }

// Traditional function name
func name(x, y MyStruct) int {
	return (x.x + y.x) * (x.y + y.y)
}

func (m MyStruct) name(y MyStruct) int {
	return (m.x + y.x) * (m.y + y.y)
}

func main() {

	m1 := MyStruct{1, 2}
	m2 := MyStruct{2, 3}
	fmt.Println(name(m1, m2))

	m1.name(m2)

}
