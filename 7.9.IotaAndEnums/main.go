//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.9: Enum in structs
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// defined custom defined type for constants
type UserLevel uint8

const (
	Beginner     UserLevel = 1
	Weak         UserLevel = 2
	Strong       UserLevel = 3
	Professional UserLevel = 4
	GodLevel     UserLevel = 5
)

const (
	Bot UserLevel = iota + 6
	User
	Moder
	Admin
)

const (
	BRT   = iota - 3
	GST   // -2
	AZOST // -1
	_     // 0
	_
	_
	MSK
)

const (
	ONE = -(1 + iota) // -1
	TWO
	THREE
	FOUR // -4
)

const (
	F   = 1
	S   = 2
	T   = iota + 2 // 4
	FOR            // 5
	FIV            // 6
)

const (
	Active, Moving, Running = iota + 1, iota + 2, iota + 5
	Passive, Stopped, Stale // 2 3 6
	//Next // error
	New   = 10 // ok
	Reset = iota
)

type Even bool

const (
	a1 = Even(iota%2 == 0) // true
	b2                     // false
	c3                     // true
	d4                     // false
)

const (
	a = string(iota + 'a') // a
	b                      // b
	c                      // c
	d                      // d
	e                      // e
)

type Season int

const (
	Winter  Season = 1 << iota
	Spring         // 2
	Summer         // 4
	Autumn         // 8
	AllYear = Winter | Spring | Summer | Autumn
)

// Entrypoint
func main() {
	// Print name of user level
	fmt.Println(Weak)     // Weak
	fmt.Println(GodLevel) // God Level

	fmt.Println(Beginner.Target()) // false
	fmt.Println(GodLevel.Target()) // true

	fmt.Println(Admin) // God

	fmt.Println(T)   // 4
	fmt.Println(FIV) // 6

	fmt.Println(Passive, Stopped, Stale)

	fmt.Println(Winter & AllYear) // 1

}

// Print user level
func (u UserLevel) String() string {
	names := [...]string{
		"Beginner",
		"Weak",
		"Strong",
		"Professional",
		"God Level",

		"Machina",
		"Human",
		"Super human",
		"God",
	}

	return names[u-1]
}

// Check is target user level
func (u UserLevel) Target() bool {
	switch u {
	case Strong, Professional, GodLevel:
		return true
	default:
		return false
	}
}
