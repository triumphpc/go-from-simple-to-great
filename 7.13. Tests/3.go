//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Correct the errors in the names of variables and / or the method in the example below
//
// func (myType MyTypeInt) Add(number int) MyTypeInt  {
// 	return myType + MyTypeInt(number)
// }
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

type MyTypeInt int

func main() {
	one := MyTypeInt(5)
	fmt.Println(one.Add(5)) // 10

}

func (m MyTypeInt) Add(n int) MyTypeInt {
	return m + MyTypeInt(n)
}
