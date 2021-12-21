//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function to count the number of each unique elements in an array
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

func main() {
	var arr = [10]int{1, 1, 3, 4, 5, 5, 3, 2, 3, 7}

	for i, c := range uniqInArr(arr[:]) {
		fmt.Println(i, " - ", c)
	}

}

// Count of value in array
func uniqInArr(a []int) (dict map[int]int) {
	dict = make(map[int]int)
	for _, v := range a {
		dict[v]++
	}
	return
}
