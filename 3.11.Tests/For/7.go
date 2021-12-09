//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. Write a program which get multiple user inputs of number
// 2. Generates an random number
// 3. If one of the entered user numbers is equal to random - display congratulations
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	n := rand.Intn(len(args))

	fmt.Println("Rand number", n)

	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
		if val < 0 {
			fmt.Println("Please enter a positive number.")
			return
		}
		if n == val {
			fmt.Println("YOU WIN!")
			return
		}
	}

	fmt.Println("YOU LOST")
}
