//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 2.21: Practice program
// Part 1. Creation of the general structure of the program
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Max enter args validator
var maxEnterArgs int = 3
var maxRate, minRate int = 36, 0

// Entrypoint
func main() {
	// Get user enter rate
	rate, volume := getUserRate()
	fmt.Println("Your rate: ", rate, " and volume: ", volume, "$")

	// Get rand number
	targetNumber := getNumber()
	fmt.Println("Roulette give a number: ", targetNumber)

	// Validate user rate
	validateRate(rate, volume)

	// Check win
	fmt.Printf("Your winnings: %.2f ", getWin(rate, volume, targetNumber))
}

// Check, if user win some money
func getWin(r int, v int, t int) float64 {
	// check odd and even
	tIsOdd := t&1 == 1
	rIsOdd := r&1 == 1
	var win float64 = 0

	// user win
	if (tIsOdd && rIsOdd) || (!tIsOdd && !rIsOdd) {
		// multiply the winnings by two
		win += float64(v) * 2
	}

	// check number
	if t == r {
		win += float64(v) * 36
	}
	return win
}

/*
Validate user rate.
r it's user enter rate
*/
func validateRate(r int, v int) {
	if r > maxRate || r < minRate {
		log.Fatal("Not valid rate")
	}

	if v <= 0 {
		log.Fatal("Not valid volume")
	}
}

// Generate rand number for roulette
func getNumber() int {
	// current seconds in unixtime
	seconds := time.Now().Unix()
	// random generator
	rand.Seed(seconds)
	// generate rand number for roulette
	// using rand package
	targetNumber := rand.Intn(37)
	// return rand number
	return targetNumber
}

// Get rate from user and rate volume
func getUserRate() (int, int) {
	// check enter from args
	var args = os.Args

	// validate count
	if len(args) == maxEnterArgs {
		rate, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		volume, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}
		return rate, volume
	}

	// enter from io interface
	// create bufio.Reader from stdio
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please, enter rate: ")
	input, err := reader.ReadString('\n') // rune new line
	if err != nil {
		log.Fatal(err)
	}
	// sanitize enter string
	input = strings.TrimSpace(input)

	// convert to int
	rate, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	// and get volume
	fmt.Print("Please, enter volume: ")
	input, err = reader.ReadString('\n') // rune new line
	if err != nil {
		log.Fatal(err)
	}
	// sanitize enter string
	input = strings.TrimSpace(input)

	// convert to int
	volume, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	return rate, volume
}
