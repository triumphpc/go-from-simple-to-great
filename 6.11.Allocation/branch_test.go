package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	//// test stack
	mainAnalysisTest()
	//
	//// test 2
	_ = stackIt2Test()

	// test 3
	main3()
}

func mainAnalysisTest() {
	_ = stackItTest()
}

//go:noinline
func stackItTest() int {
	y := 4
	return y * 4
}

//go:noinline
func stackIt2Test() *int {
	y := 2
	res := y * 2
	return &res
}

func main3() {
	y := 2
	_ = stackIt3Test(&y) // pass y down the stack as a pointer
}

//go:noinline
func stackIt3Test(y *int) int {
	res := *y * 2
	return res
}
