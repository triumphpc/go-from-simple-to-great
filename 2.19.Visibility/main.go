// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.19: Visibility blocks
//
// @author Alex Versus 2021
// www.alexversus.com

// package block
package main

import "fmt"

var packageAreaVar = "package"

// function block
func main() {
	var functionAreaVar = "some var"

	// condition block
	if true {
		var conditionAreaVar = "condition var"
		fmt.Println(conditionAreaVar)
		fmt.Println(packageAreaVar)
		fmt.Println(functionAreaVar)

		// condition block
		if false {
			fmt.Println(conditionAreaVar)
			fmt.Println(packageAreaVar)
			fmt.Println(functionAreaVar)

			var conditionInsideAreaVar = "second condition var"
			fmt.Println(conditionInsideAreaVar)
		}
		//fmt.Println(conditionInsideAreaVar)
	}

	fmt.Println(packageAreaVar)
	fmt.Println(functionAreaVar)
	//fmt.Println(conditionAreaVar)
	//fmt.Println(conditionInsideAreaVar)
}
