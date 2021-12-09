//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 3.8: Goto statement
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint
func main() {
	fmt.Println("Ready")

	// jump to label
	goto StepOne

	fmt.Println("Steady")

StepOne:
	fmt.Println("It's step one")

	fmt.Println("It's step two")

	goto StepTwo

	fmt.Println("It's step three")

StepTwo:
	fmt.Println("It's step four")

	// variable name
	var StepOne = "some string"
	fmt.Println("It's", StepOne)

	// labels in for statement
FirstLoop:
	for i := 0; i < 5; i++ {
	SecondLoop:
		for j := 0; j < 5; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			break SecondLoop
		}
		continue FirstLoop
		fmt.Println("Never print")
	}

	// labels in switch statement
	var i = 0
	fmt.Println("Switch example")
SwitchSt:
	switch i%2 == 0 {
	case true:
		i++
		fmt.Println(1)
		for j := 0; j < 10; j++ {
			// exit from switch
			break SwitchSt
		}
		fallthrough
	case false:
		fmt.Println(0)
	}

	// example with goto
	i = 0
SwitchSec:
	switch i%2 == 0 {
	case true:
		i++
		fmt.Println(1)
		for j := 0; j < 10; j++ {
			// exit from switch
			goto SwitchSec
		}
		fallthrough
	case false:
		goto SwitchEnd
		fmt.Println(0)
	}
SwitchEnd:
	fmt.Println("Yep")

}
