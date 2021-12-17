//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Lesson 5.7: Practice program
//
//Part 4.
// 1. Display interactive information on the possibilities of playing roulette.
// 2. Add user profile data input mode
// 3. Add the ability to place bets on the color red / black
// 4. Add the ability to enter multiple rates
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
type userProfile struct {
	firstName, secondName string
	age                   uint8
	phone                 uint
}

// constants
const (
	// the same value
	maxInputCount uint8 = 3
	// max and min user rate
	maxRate, minRate uint8 = 36, 0
	// odd and even values
	odd, even uint8 = 101, 102
)

// Entrypoint
func main() {
	var (
		userBets   map[string]Cent
		inputCount uint8 = 0
	)

	// Get user profile
	up := getPersonalData()

	fmt.Println("Welcome to the Roulette game. \n" +
		"You can place bets as follows: \n" +
		"<0 ... 36> <volume> - to place a bet on a numeric field \n" +
		"<odd / even> <volume> - to place a bet on odd / even fields \n" +
		"<red / black> <volume> - to place a bet bets on red / black fields " +
		"<q> - exit the bet entry mode and calculate the results.")

	for {
		ub, err := getUserRate()
		if err != nil {
			inputCount++
			if inputCount >= maxInputCount {
				os.Exit(0)
			}
		} else {
			userBets = ub
			break
		}
		fmt.Println("Incorrect enter. Please enter again. You have ", maxInputCount-inputCount, "attempts")
	}

	fmt.Println("Your rates:")
	for bet, vol := range userBets {
		fmt.Println(bet, "-", vol)
	}

	// Get rand number
	targetNumber := getNumber()
	fmt.Println("Roulette give a number: ", targetNumber)

	// Check win
	w := getWin(userBets, targetNumber)
	if w > 0 {
		fmt.Printf("Congratulations. You won: %.2f$ \n", toDollar(w))
		fmt.Println("Your personal data:")
		fmt.Println(up.firstName, up.secondName)
		fmt.Println("Age:", up.age, "Phone:", up.phone)
	} else {
		fmt.Printf("Sorry, you lost\n")
	}
}

// Get personal data
func getPersonalData() userProfile {
	reader := bufio.NewReader(os.Stdin)
	var up userProfile
	for {
		fmt.Print("Enter your first name:")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			continue
		}
		up.firstName = input
		break
	}

	for {
		fmt.Print("Enter your second name:")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			continue
		}
		up.secondName = input
		break
	}

	for {
		fmt.Print("Enter your age:")
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		// sanitize enter string
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input) // is numeric
		if err != nil {
			continue
		}
		up.age = uint8(num)
		break
	}

	for {
		fmt.Print("Enter your phone:")
		input, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		// sanitize enter string
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input) // is numeric
		if err != nil {
			continue
		}
		up.phone = uint(num)
		break
	}
	return up
}

// Check, if user win some money
func getWin(userBets map[string]Cent, t uint8) Cent {
	// check odd and even
	tIsOdd := t&1 == 1
	// red fields
	rf := [18]uint8{1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36}
	var win Cent = 0
	tIsRed := false
	for _, a := range rf {
		if t == a {
			tIsRed = true
			break
		}
	}

	for r, v := range userBets {
		// check rate
		num, err := strconv.Atoi(r) // is numeric

		if err != nil {
			switch r {
			case "odd":
				if tIsOdd {
					win += v * 2
				}
			case "even":
				if !tIsOdd {
					win += v * 2
				}
			case "red":
				if tIsRed {
					win += v * 2
				}
			case "black":
				if !tIsRed {
					win += v * 2
				}
			}
		} else {
			// check number
			if t == uint8(num) {
				win += v * 36
			}
		}
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

/*
Generate rand number for roulette
*/
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

// Get rate from user and rate volume
func getUserRate() (map[string]Cent, error) {
	// enter from io interface
	reader := bufio.NewReader(os.Stdin)
	var ur = make(map[string]Cent)

	for {
		fmt.Print("Please, enter bet: ")

		input, err := reader.ReadString('\n') // rune new line
		if err != nil {
			return nil, err
		}
		// sanitize enter string
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input) // is numeric
		isNum := false

		if err != nil {
			switch input {
			case "odd":
			case "even":
			case "red":
			case "black":
				ur[input] = 0
			case "q":
				return ur, nil

			default:
				return nil, err
			}
		} else {
			isNum = true
			if _, ok := ur[input]; !ok {
				ur[input] = 0
			}
		}

		// and get volume
		fmt.Print("Please, enter volume: ")
		inputVolume, err := reader.ReadString('\n') // rune new line
		if err != nil {
			return nil, err
		}

		// convert to float64
		volume, err := strconv.ParseFloat(strings.TrimSpace(inputVolume), 64)
		if err != nil {
			return nil, err
		}

		// Validate user rate
		vic := toCent(Dollar(volume))
		if isNum {
			validateRate(uint8(num), vic)
		}

		ur[input] += vic
	}
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
