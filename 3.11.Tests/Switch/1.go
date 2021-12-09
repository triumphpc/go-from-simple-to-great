//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// 1. The program accepts the current age as input
// 2. If the age is less than 1 - the program displays "Breast"
// 3. If the age is from 1 to 3 - the program displays "Early childhood"
// 4. If the age is from 3 to 7 - the program displays "Preschool"
// 5 If the age is from 7 to 12 - the program displays "Junior school"
// 6. If the age is from 12 to 18 - the program displays "Teenage"
// 7. If the age is from 18 to 21 - the program displays "Junior"
// 8. If the age is from 21 to 60 - the program displays "Mature"
// 9. If the age is from 60 to 75 - the program displays "Elderly"
// 10. If the age is from 75 - "Senile"
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please enter your age.")
		return
	}

	a, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Try again.")
		return
	}

	var d string

	switch r := a; {
	case r < 1:
		d = "Breast"
	case r >= 1 && r < 3:
		d = "Early childhood"
	case r >= 3 && r < 7:
		d = "Preschool"
	case r >= 7 && r < 12:
		d = "Junior school"
	case r >= 12 && r < 18:
		d = "Teenage"
	case r >= 18 && r < 21:
		d = "Junior"
	case r >= 21 && r < 60:
		d = "Mature"
	case r >= 60 && r < 75:
		d = "Elderly"
	default:
		d = "Senile"
	}

	fmt.Printf("%.2f is %s\n", a, d)

}
