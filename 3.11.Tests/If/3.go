//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Simplify the code below using if statements
//
// ------------------------
// maxSpeed, isRace, currentSpeed := 250, false, 200
//
// if currentSpeed >= 50 {
// 	if currentSpeed >= 100 {
// 		if currentSpeed > maxSpeed {
// 			isRace = true
// 		}
// 	}
// }
//
// if currentSpeed > maxSpeed {
// 	fmt.Println("This is race.")
// } else if !(isRace == false) {
// 	fmt.Println("This is race.")
// } else {
// 	fmt.Println("This is road speed.")
// }
// ------------------------
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	maxSpeed, _, currentSpeed := 250, false, 200

	if currentSpeed > maxSpeed {
		fmt.Println("This is race.")
	} else {
		fmt.Println("This is road speed.")
	}
}
