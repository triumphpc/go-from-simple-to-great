//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 3.3: Nil-value and error handling
//
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
	fmt.Println("Your age: ", userAge)

	// example work with nil
	err = getNil()

	if err != nil {
		log.Fatal("Some problem here")
	}
	fmt.Println("And we continue")

	// example of functions without return nil
	b := strconv.CanBackquote("Example")
	fmt.Println(b)

	s, err := strconv.Atoi("fdf")
	fmt.Println("It's val:", s)
	fmt.Println("It's err:", err)

}

func getNil() error {
	return nil

}
