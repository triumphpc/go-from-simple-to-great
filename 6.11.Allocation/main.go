//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//  Lesson 6.11: Allocation
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"bytes"
	"fmt"
)

type Buffer struct {
	buffer bytes.Buffer
}

type StructType struct {
	one, two, three int
}

// Entrypoint
func main() {
	p := new(Buffer) //&{{[] 0 0}}

	var p2 Buffer // {{[] 0 0}}

	fmt.Println(p, p2)

	// no allocation
	var i int
	i = 1
	i = 2
	fmt.Println(i)

	j := 1
	j = 2
	fmt.Println(j)

	z2 := new(StructType)
	fmt.Println(z2) // &{0 0 0}

	// test stack
	mainAnalysis()

	// test 2
	_ = stackIt2()

	// test 3
	y := 2
	_ = stackIt3(&y) // pass y down the stack as a pointer

}

func mainAnalysis() {
	_ = stackIt()
}

//go:noinline
func stackIt() int {
	y := 4
	return y * 4
}

//go:noinline
func stackIt2() *int {
	y := 2
	res := y * 2
	return &res
}

//go:noinline
func stackIt3(y *int) int {
	res := *y * 2
	return res
}
