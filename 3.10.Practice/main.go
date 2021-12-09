//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 3.10: Practice program
//
// Part 2.
// 1. Add three attempts to enter
// 2. Print a congratulations from the program if the user's bet really wins.
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
var (
	maxEnterArgs     int  = 3
	maxRate, minRate int  = 36, 0
	inputCount       int8 = 0
	maxInputCount    int8 = 3
)

// Entrypoint
func main() {
	// Get user enter rate
	var (
		rate   = 0
		volume = 0
	)

	for {
		r, v, err := getUserRate(inputCount)
		if err != nil {
			if r > 0 {
				// reset counter
				inputCount = 0
			}
			inputCount++
			if inputCount >= maxInputCount {
				os.Exit(0)
			}
		} else {
			rate = r
			volume = v
			break
		}
		fmt.Println("Incorrect enter. Please enter again. You have ", maxInputCount-inputCount, "attempts")
	}

	fmt.Println("Your rate: ", rate, " and volume: ", volume, "$")

	// Get rand number
	targetNumber := getNumber()
	fmt.Println("Roulette give a number: ", targetNumber)

	// Validate user rate
	validateRate(rate, volume)

	// Check win
	w := getWin(rate, volume, targetNumber)
	if w > 0 {
		fmt.Printf("Congratulations. You won: %.2f ", w)
	} else {
		fmt.Printf("Sorry, you lost")
	}
}

/*
Check, if user win some money
*/
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

/*
Generate rand number for roulette
*/
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

/*
Get rate from user and rate volume
*/
func getUserRate(inputCount int8) (int, int, error) {
	// check enter from args
	var args = os.Args

	// validate count
	if inputCount == 0 && len(args) == maxEnterArgs {
		rate, err := strconv.Atoi(args[1])
		if err != nil {
			return 0, 0, err
		}
		volume, err := strconv.Atoi(args[2])
		if err != nil {
			return 0, 0, err
		}
		return rate, volume, nil
	}

	// enter from io interface
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please, enter rate: ")
	input, err := reader.ReadString('\n') // rune new line
	if err != nil {
		return 0, 0, err
	}
	// sanitize enter string
	input = strings.TrimSpace(input)

	// convert to int
	rate, err := strconv.Atoi(input)
	if err != nil {
		return 0, 0, err
	}

	// and get volume
	fmt.Print("Please, enter volume: ")
	input, err = reader.ReadString('\n') // rune new line
	if err != nil {
		return 0, 0, err
	}
	// sanitize enter string
	input = strings.TrimSpace(input)

	// convert to int
	volume, err := strconv.Atoi(input)
	if err != nil {
		return 0, 0, err
	}

	return rate, volume, nil
}
