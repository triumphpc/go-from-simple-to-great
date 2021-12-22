//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Lesson 8.8: Practice program. Error handling
//
// Part 7.
// 1. Add error handling, apply different strategies for dealing with errors.
// 2. Add panic to our program if the user has finished trying to enter rates.
// 3. Add a recovery mode in case of further rates input.
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
	"reflect"
	"strconv"
	"strings"
	"time"
)

// defined types
type Dollar float64
type Cent float64

// Type for full name of user
type FullInfoType func(string, Cent) string

type userProfile struct {
	firstName  string `title:"User Name"`
	secondName string `title:"User Second Name"`
	age        uint8  `title:"User Age"`
	phone      uint   `title:"User phone"`
	info       FullInfoType
	bets       map[string]Cent
	level      UserLevel
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

// defined custom defined type for constants
type UserLevel uint8

// Counter for attempts
var counter uint8

const (
	Beginner UserLevel = iota
	Strong
	Professional
)

// Entrypoint
func main() {
	// Get user profile
	up := personalData()

	// Print enter information
	printWelcome()

	// Get bets from user
	up.SetBets(userBets())

	// Get rand number
	tn := number()
	// Check win
	w := win(up.bets, tn)
	// Print result
	printResult(w, up)
}

// Run logic for get user bets
func userBets() map[string]Cent {
	return bets()
}

func recoverUserBets() {
	p := recover()
	err, ok := p.(error)

	if ok {
		fmt.Println(fmt.Errorf("get bets errors: %s", err))
		fmt.Print("Do you want continue enter bets? ")
		input := getUserInput()

		if input == "y" {
			userBets()
		} else {
			log.Fatalf("Error %v\n", err)
		}
	}
}

// Print result for user
func printResult(w Cent, up userProfile) {
	if w > 0 {
		fmt.Println("Your personal data:")

		pr := reflect.TypeOf(up)
		f, _ := pr.FieldByName("firstName")
		fmt.Println(f.Tag.Get("title"), up.Name())

		f, _ = pr.FieldByName("secondName")
		fmt.Println(f.Tag.Get("title"), up.SecondName())

		f, _ = pr.FieldByName("age")
		fmt.Println(f.Tag.Get("title"), up.Age())

		fmt.Println("Your level is:", up.Level())

		f, _ = pr.FieldByName("phone")
		fmt.Println(f.Tag.Get("title"), up.Phone())

		fullName := up.Info()
		fmt.Println(fullName(up.Name(), w))

		if up.Level() == Beginner {
			fmt.Println("Sorry, but you can't take your prize, because you only ", Beginner)
		}

	} else {
		fmt.Printf("Sorry, you lost\n")
	}
}

// Print information data
func printWelcome() {
	fmt.Println("Welcome to the Roulette game. \n" +
		"You can place bets as follows: \n" +
		"<0 ... 36> <volume> - to place a bet on a numeric field \n" +
		"<odd / even> <volume> - to place a bet on odd / even fields \n" +
		"<red / black> <volume> - to place a bet bets on red / black fields " +
		"<q> - exit the bet entry mode and calculate the results.")

}

// Get bets from user
func bets() map[string]Cent {
	defer recoverUserBets()

	ub, err := userRate()

	if err != nil {
		counter++
		if counter >= maxInputCount {
			panic(fmt.Errorf("you have run out of input attempts %d \n", counter))
		}
		fmt.Println("Incorrect enter. Please enter again. You have ", maxInputCount-counter, "attempts")
		return bets()
	}

	// it's call when we out from self function
	defer printBets(ub)

	return ub
}

// Print all user bets
func printBets(ub map[string]Cent) {
	fmt.Println("Your rates:")
	for bet, vol := range ub {
		fmt.Println(bet, "-", vol)
	}
}

// Get personal data
func personalData() (up userProfile) {
	fmt.Print("Enter your first name:")
	input := getUserInput()
	up.setName(input)

	fmt.Print("Enter your second name:")
	input = getUserInput()
	up.setSecondName(input)

	for {
		fmt.Print("Enter your age:")
		input = getUserInput()

		num, err := strconv.Atoi(input) // is numeric
		if err != nil {
			continue
		}

		up.setAge(uint8(num))
		break
	}

	for {
		fmt.Print("Enter your phone:")
		input = getUserInput()

		num, err := strconv.Atoi(input) // is numeric
		if err != nil {
			continue
		}
		up.setPhone(uint(num))
		break
	}

	info := func(fn string, w Cent) string {
		return fmt.Sprintf("Congratulations. You won: %.2f$ \n", toDollar(w))
	}
	up.setInfo(info)
	return
}

// Get input user strategy
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return getUserInput()
	}

	// Sanitize user input
	input = strings.TrimSpace(input)
	return input
}

/*
Check, if user win some money
*/
func win(userBets map[string]Cent, t uint8) (w Cent) {
	// check odd and even
	tIsOdd := t&1 == 1
	// red fields
	rf := [18]uint8{1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36}
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
					w += v * 2
				}
			case "even":
				if !tIsOdd {
					w += v * 2
				}
			case "red":
				if tIsRed {
					w += v * 2
				}
			case "black":
				if !tIsRed {
					w += v * 2
				}
			}
		} else {
			// check number
			if t == uint8(num) {
				w += v * 36
			}
		}
	}
	return
}

/*
Validate user rate.
it's user enter rate
*/
func validateRate(r uint8, v Cent) {
	if r > maxRate || r < minRate {
		if r != odd && r != even {
			log.Fatal("Not valid bet")
		}
	}

	if v <= 0 {
		log.Fatal("Not valid volume")
	}
}

// Generate rand number for roulette
func number() uint8 {
	// current seconds in unixtime
	seconds := time.Now().Unix()
	// random generator
	rand.Seed(seconds)
	// generate rand number for roulette
	// using rand package
	targetNumber := rand.Intn(37)

	fmt.Println("Roulette give a number: ", targetNumber)
	// return rand number
	return uint8(targetNumber)
}

/*
Get rate from user and rate volume
*/
func userRate() (ur map[string]Cent, err error) {
	// enter from io interface
	reader := bufio.NewReader(os.Stdin)
	// declare map alloc
	ur = make(map[string]Cent)

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

// Set name method for user profile
func (p *userProfile) setName(n string) {
	p.firstName = n
}

// Set second name method for user profile
func (p *userProfile) setSecondName(n string) {
	p.secondName = n
}

// Set age method for user profile
func (p *userProfile) setAge(a uint8) {
	p.age = a

	if a < 18 {
		p.level = Beginner
	} else if a < 60 {
		p.level = Strong
	} else {
		p.level = Professional
	}
}

// Set phone method for user profile
func (p *userProfile) setPhone(ph uint) {
	p.phone = ph
}

// Set info anonymous func method for user profile
func (p *userProfile) setInfo(i FullInfoType) {
	p.info = i
}

// Get name method for user profile
func (p userProfile) Name() string {
	return p.firstName
}

// Get second name method for user profile
func (p userProfile) SecondName() string {
	return p.secondName
}

// Get age method for user profile
func (p userProfile) Age() uint8 {
	return p.age
}

// Get phone method for user profile
func (p userProfile) Phone() uint {
	return p.phone
}

// Get func for take full info method for user profile
func (p userProfile) Info() FullInfoType {
	return p.info
}

// Set user bet
func (p *userProfile) SetBets(ur map[string]Cent) {
	p.bets = ur
}

// Get user bets
func (p userProfile) Bets() map[string]Cent {
	return p.bets
}

// Print user level
func (u UserLevel) String() string {
	names := [...]string{
		"Beginner",
		"Strong",
		"Professional",
	}

	return names[u]
}

// Get user level
func (p userProfile) Level() UserLevel {
	return p.level
}
