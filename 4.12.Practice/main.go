//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 4.12: Practice program

// Part 3.
// 1. Let's take out all constant values into constants and we will use constants in the program.
// 2. Let's add the ability to place a bet not only with an integer volume number, but also a bet volume in the form of a floating point number.
// To get the volume as a floating point number, we use a string to float conversion.
// 3. Let's update the logic by calculating the gain taking into account the volume in the form of a floating point number.
// 4. We declare the user's bets in the form of a special definable type.
// 5. We will enable the user to place bets on even odd.
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

// defined types
type Dollar float64
type Cent float64

// constants
const (
	// maximum allowed amount of arguments
	maxEnterArgs uint8 = 3
	// the same value
	maxInputCount
	// max and min user rate
	maxRate, minRate uint8 = 36, 0
	// odd and even values
	odd, even uint8 = 101, 102
)

// Entrypoint
func main() {
	var (
		volume           Cent  = 0.0
		inputCount, rate uint8 = 0, 0
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

	fmt.Printf("Your rate: %d and volume: %.0f cents\n", rate, volume)

	// Validate user rate
	validateRate(rate, volume)

	// Get rand number
	targetNumber := getNumber()
	fmt.Println("Roulette give a number: ", targetNumber)

	// Check win
	w := getWin(rate, volume, targetNumber)
	if w > 0 {
		fmt.Printf("Congratulations. You won: %.2f$ \n", toDollar(w))
	} else {
		fmt.Printf("Sorry, you lost\n")
	}
}

/*
Check, if user win some money
*/
func getWin(r uint8, v Cent, t uint8) Cent {
	// check odd and even
	tIsOdd := t&1 == 1
	rIsOdd := r&1 == 1
	var win Cent = 0

	// user win
	if (tIsOdd && rIsOdd) || (!tIsOdd && !rIsOdd) {
		// multiply the winnings by two
		win += v * 2
	}

	// check number
	if t == r {
		win += v * 36
	}
	return win
}

/*
Validate user rate.
it's user enter rate
*/
func validateRate(r uint8, v Cent) {
	if r > maxRate || r < minRate {
		if r != odd && r != even {
			log.Fatal("Not valid rate")
		}
	}

	if v <= 0 {
		log.Fatal("Not valid volume")
	}
}

// Generate rand number for roulette
func getNumber() uint8 {
	// current seconds in unixtime
	seconds := time.Now().Unix()
	// random generator
	rand.Seed(seconds)
	// generate rand number for roulette
	// using rand package
	targetNumber := rand.Intn(37)
	// return rand number
	return uint8(targetNumber)
}

/*
Get rate from user and rate volume
*/
func getUserRate(inputCount uint8) (uint8, Cent, error) {
	// check enter from args
	var args = os.Args

	// validate count
	if inputCount == 0 && uint8(len(args)) == maxEnterArgs {
		rate, err := getRate()
		if err != nil {
			return 0, 0, err
		}
		volume, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return 0, 0, err
		}
		return rate, toCent(Dollar(volume)), nil
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

	var rate uint8

	if input == "o" {
		rate = odd
	} else if input == "e" {
		rate = even
	} else {
		r, err := strconv.Atoi(input)
		if err != nil {
			return 0, 0, err
		}
		rate = uint8(r)
	}
	// and get volume
	fmt.Print("Please, enter volume: ")
	input, err = reader.ReadString('\n') // rune new line
	if err != nil {
		return 0, 0, err
	}
	// sanitize enter string
	input = strings.TrimSpace(input)

	// convert to float64
	volume, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, 0, err
	}

	return rate, toCent(Dollar(volume)), nil
}

/**
Get rate from user
*/
func getRate() (uint8, error) {
	var args = os.Args

	if args[1] == "o" {
		// system value
		return odd, nil
	}
	if args[1] == "e" {
		// system value
		return even, nil
	}

	rate, err := strconv.Atoi(args[1])
	return uint8(rate), err
}

/**
Convert Dollar to Cent
*/
func toCent(d Dollar) Cent {
	return Cent(d * 100)
}

/**
Convert Cent to Dollar
*/
func toDollar(c Cent) Dollar {
	return Dollar(c / 100)
}
