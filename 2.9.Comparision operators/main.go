// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.9: Comparison operators
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint for program
func main() {
	// equal
	var first int = 1
	var second int = 2
	var third string = "not"
	var forth string = "equal"

	fmt.Printf("|%15s|%15s|%15s|%15s|\n", "name", "operator", "example", "result")
	fmt.Println("|-----------------------------------")
	fmt.Printf("|%15s|%15s|%15s|%15t|\n", "equal", "==", "1==2", first == second)
	fmt.Printf("|%15s|%15s|%15s|%15t|\n", "not equal", "!=", "'not'!='equal'", third != forth)
	fmt.Printf("|%15s|%15s|%15s|%15t|\n", "less", "<", "2 < 1", second < first)
	fmt.Printf("|%15s|%15s|%15s|%15t|\n", "less or equal", "<=", "2 <= 1", 4 <= 5)
	fmt.Printf("|%15s|%15s|%15s|%15t|\n", "greater", ">", "\"a\" > \"c\"", "a" > "c")
	fmt.Printf("|%15s|%15s|%15s|%15t|\n", "greater or equal", ">=", "\"z\" > \"y\"", "z" >= "y")
}
