//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. What land for constant secondConst ?
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

const (
	firstConst  string = "That first constant value"
	secondConst        = firstConst + " ❎ "
)

func main() {
	fmt.Println(firstConst, len(firstConst))
	fmt.Println(secondConst, len(secondConst))

}
