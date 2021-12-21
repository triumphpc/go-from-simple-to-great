//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.2: Method declaration
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Custom type for method
type typeForMethod string

// Entrypoint
func main() {

	v := typeForMethod("value string")

	// v - method recipient
	// methodName -  method name
	v.methodName()

	// call with params
	v2 := typeForMethod("new string")
	v2.methodWithParams(5, true)

	// return params
	v3 := typeForMethod("string")
	fmt.Println(v3.returnMethod()) // 6 123

}

// Method declaration
// m - recipient parameter name
// typeForMethod - recipient parameter type
// methodName - method name
func (m typeForMethod) methodName() {
	fmt.Println("Call method", m)
}

// Method with additional params
func (m typeForMethod) methodWithParams(n int, f bool) {
	fmt.Println("Call method with recipient parameter", m, "and params", n, f)

}

// Return several params
func (m typeForMethod) returnMethod() (int, int) {
	return len(m), 123
}
