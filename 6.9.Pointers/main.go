//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.9: Pointers
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"reflect"
)

// nil value
type NilSt struct {
	name string
}

// Entrypoint
func main() {
	a := 5
	fmt.Println(a)
	fmt.Println(&a) // print address

	x := 3
	a2 := &x
	fmt.Println(a2)  // address
	fmt.Println(*a2) // 3
	*a2 = 5
	fmt.Println(x)                   // 5
	fmt.Println(reflect.TypeOf(*a2)) // int

	// point on int
	var myInt int
	var pointInt *int
	pointInt = &myInt
	fmt.Println(pointInt)

	*pointInt = 5
	fmt.Println(myInt) // 5

	var v uint = 1
	increase(&v)
	increase(&v)
	fmt.Println(v) // 3

	// use point in struct example
	type Employee struct {
		name string
		id   int
	}

	emp := Employee{"Alex", 123}
	p := &emp
	fmt.Println(p)  // &{Alex 123}
	fmt.Println(*p) // {Alex 123}
	p.name = "New name"
	fmt.Println(emp) // {New name 123}

	// pass by value
	y := 1
	someFunc(&y)
	fmt.Println(y) // 5

	var st *NilSt
	fmt.Println(st) // nil

	//change(st) // panic

	st = &NilSt{name: "Hello"}
	fmt.Printf("1) %+v\n", st)
	change(st)
	fmt.Printf("3) %+v\n", st) // 3) &{name:New value}

}

func increase(v *uint) uint {
	*v++
	return *v
}

// *int means you *must* pass a *int (pointer to int), NOT just an int!
func someFunc(x *int) {
	*x = 5 // Whatever variable caller passed in will now be 5
	y := 9
	x = &y // has no impact on the caller because we overwrote the pointer value!
}

func change(st *NilSt) {
	if st == nil {
		fmt.Println("Can't change")
		return
	}
	st.name = "New value"
	fmt.Printf("2) %+v\n", st)
}
