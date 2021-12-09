// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lessons 3.5: Switch operator
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint
func main() {
	switch getString() {
	case "one":
		fmt.Println("It's one person")
		one := "this is one"
		fmt.Println(one)
	case "second", "2", "two":
		fmt.Println("It's second person")
		//fmt.Println(one) // you can't use it
	default:
		fmt.Println("Year baby")
	}

	// from int
	i := getNumber()
	switch i {
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("1")
	}

	// as expressions
	number := getString()
	switch {
	case number == "one" || number == "1":
		fmt.Println("It's one person")
	case number == "second":
		fmt.Println("It's second person")
	default:
		fmt.Println("Year baby")
	}

	// Simple statement

	switch num := 10; {
	case num > 10:
		fmt.Println("It's more than 10")

	case num < 10:
		fmt.Println("It's less than 10")
	default:
		fmt.Println("It's equal 10")
	}

	// Fallthrough example
	dishes := []string{"hamburger", "pizza", "soup", "banana"}

	for _, dish := range dishes {
		switch dish {
		case "hamburger":
			fmt.Println(dish, "is my favorite!")
			fallthrough
		case "pizza", "soup":
			fmt.Println(dish, "is great yeea!")
		default:
			fmt.Println("I've never tried", dish, "before")
		}
	}

}

/*
Get number for switch
*/
func getString() string {
	return "two"
}

func getNumber() int {
	return 1

}
