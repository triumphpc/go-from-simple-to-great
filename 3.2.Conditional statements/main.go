// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 3.2: Conditional statements
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Entrypoint
func main() {
	// for read enter from keyboard
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please place your age: ")
	userAge, _ := reader.ReadString('\n')

	userAge = strings.TrimSpace(userAge)
	fmt.Println("Your age: ", userAge)
}
