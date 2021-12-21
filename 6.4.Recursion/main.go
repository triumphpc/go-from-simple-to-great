//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.4: Recursion
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint
func main() {
	fmt.Println(fact(5)) // 120

	// passing 0 as a parameter
	answer1 := factorialCalc(0)
	fmt.Println(answer1)

	// passing a positive integer
	answer2 := factorialCalc(5)
	fmt.Println(answer2)

	// passing a negative integer
	// prints an error message
	// with a return value of -1
	answer3 := factorialCalc(-1)
	fmt.Println(answer3)

	// passing a positive integer
	answer4 := factorialCalc(7)
	fmt.Println(answer4)

	// indirect recursion run
	printOne(8)

	// tail recursion
	tailCall(3)

	// head recursion
	headCall(3)

	// declaring anonymous function
	// that takes an integer value
	var anonFunc func(int)

	// defining an anonymous
	// function that prints
	// numbers from n to 1
	anonFunc = func(n int) {
		// base case
		if n == 0 {
			return
		} else {
			fmt.Println(n)

			// calling anonymous
			// function recursively
			anonFunc(n - 1)
		}
	}

	// call to anonymous
	// recursive function
	anonFunc(3)

}

// recursion example
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// recursive function for
// calculating a factorial
// of a positive integer
func factorialCalc(n int) int {

	// this is the base condition
	// if number is 0 or 1 the
	// function will return 1
	if n == 0 || n == 1 {
		return 1
	}

	// if negative argument is
	// given, it prints error
	// message and returns -1
	if n < 0 {
		fmt.Println("Invalid number")
		return -1
	}

	// recursive call to itself
	// with argument decremented
	// by 1 integer so that it
	// eventually reaches the base case
	return n * factorialCalc(n-1)
}

// recursive function for
// printing all numbers
// up to the number n
func printOne(n int) {

	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In first function:", n)
		// call to the second function
		// which calls this first
		// function indirectly
		printTwo(n - 1)
	}
}

func printTwo(n int) {

	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In second function:", n)
		// call to the first function
		printOne(n - 1)
	}
}

// tail recursive function
// to print all numbers
// from n to 1
func tailCall(n int) {

	// if number is still
	// positive, print it
	// and call the function
	// with decremented value
	if n > 0 {
		fmt.Println(n)

		// last statement in
		// the recursive function
		// tail recursive call
		tailCall(n - 1)
	}
}

// head recursive function
// to print all numbers
// from 1 to n
func headCall(n int) {

	// if number is still
	// less than n, call
	// the function
	// with decremented value
	if n > 0 {

		// first statement in
		// the function
		headCall(n - 1)

		// printing is done at
		// returning time
		fmt.Println(n)
	}
}
