// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Change the expressions to produce the expected outputs:
//
// OUTPUT:
// print 117
// 17 + 15 - 5 * 10
//
// print 2
// 2 / 0.5 + 11 - 10.5
//
// print 1
// 7 - 30 + 30 / 3 * 10 + 4
//
// print 10
// 10 / 11 * 10 + 1 % 11
//
// print 4 float
// 14 / 7 / 2
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	//17 + 15 - 5 * 10
	fmt.Println(17 + ((15 - 5) * 10)) // 117

	//2 / 0.5 + 11 - 10.5
	fmt.Println(2 / (0.5 + (11 - 10.5))) // 2

	//7 - 30 + 30 / 3 * 10 + 4
	fmt.Println(7 - (((30 + 30) / (3 * 10)) + 4))

	//10 / 11 * 10 + 1 % 11
	fmt.Println(10 / (((11 * 10) + 1) % 11))

	//fmt.Println(18 / 6 / 3)
	fmt.Println(14 / (7. / 2)) // 4 in float

}
