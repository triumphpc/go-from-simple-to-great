//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.7 Value-function and anonymous
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"sort"
)

// Entrypoint
func main() {

	// value function
	v := value(5)
	fmt.Println(v)

	// anonymous function
	f := anon(5)
	fmt.Println(f()) // 6
	fmt.Println(f()) // 7

	// example use
	example()

	p, n := close(), close()
	for i := 0; i < 10; i++ {
		fmt.Println(p(i), n(-2*i))
	}

}

func value(n int) int {
	return n + 5
}

func anon(n int) func() int {
	var x int
	return func() int {
		x++
		return x + n
	}
}

func example() {
	var tree = map[string][]string{
		"one": {
			"one_1",
			"one_2",
			"one_3",
		},
		"two": {
			"two_1",
			"two_2",
		},
	}

	for _, l := range sortArr(tree) {
		fmt.Println(l)
	}
}

// Sort values from tree
func sortArr(m map[string][]string) []string {
	var ord []string
	seen := make(map[string]bool)

	// declare
	var visitAll func(items []string)

	// implement
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				ord = append(ord, item)
			}
		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	visitAll(keys)
	return ord
}

// Return closure
func close() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
