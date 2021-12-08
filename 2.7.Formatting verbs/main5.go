// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.7:Formatting verbs width control
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint for program
func main() {
	fmt.Printf("|%10s|%10s|\n", "Employee", "Salary")
	fmt.Println("|---------------------|")
	fmt.Printf("|%10s|%10.2f|\n", "Alex", 123.4455)
	fmt.Printf("|%10s|%10.2f|\n", "Mike", 123.4455)
	fmt.Printf("|%10s|%10.2f|\n", "Freddy", 123.4455)
	fmt.Printf("|%10s|    %.2f|\n", "Arnold", 123.4455)
}
