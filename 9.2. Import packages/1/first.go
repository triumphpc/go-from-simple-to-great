package main

import "fmt"

func main() {
	integers := []int{1, 2, 3, 4, 5}
	fmt.Println(CalcSum(integers))

}

// Calc sum from slice of ints
func CalcSum(n []int) (s int) {
	for _, v := range n {
		s += v
	}
	return
}
