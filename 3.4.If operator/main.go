//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 3.4: If operator
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Entrypoint
func main() {
	// for read enter from keyboard
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please place your age: ")
	userAge, err := reader.ReadString('\n')

	// check on errors that err does not contain errors
	if err != nil {
		log.Fatal(err)
	}

	userAge = strings.TrimSpace(userAge)
	// convert to int

	var currentAge float64
	if c, err := strconv.ParseFloat(userAge, 64); err == nil {
		currentAge = c
	} else {
		log.Fatal(err)
	}

	// it's float64
	fmt.Printf("%T\n", currentAge)

	fmt.Println("Your age: ", currentAge)

	// conditional of age
	if currentAge == 0 || currentAge < 18 {
		log.Fatal("You can't use this program")
	} else if currentAge > 60 {
		log.Fatal("You very old, man")
	} else {
		fmt.Println("Ok, lets start")
	}
}
