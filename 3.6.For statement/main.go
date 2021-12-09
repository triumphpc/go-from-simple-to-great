// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 3.6: For operator
//
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import "fmt"

// Entrypoint
func main() {
	// ForClause
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration:", i)
	}

	for i := 15; i > 10; i-- {
		fmt.Println("Iteration:", i)
	}

	// while loop
	j := 14
	for j >= 10 {
		fmt.Println("while loop Iter:", j)
		j--
	}

	// break example
	k := 100
	for k > 10 {
		k--
		fmt.Println("has loop Iter:", k)

		if k == 95 {
			fmt.Println("break...")
			break
		}
	}

	// continue example
	for k > 95 {
		k--
		fmt.Println("has loop Iter:", k)

		if k == 98 {
			fmt.Println("continue...")
			break
		}
	}

	// endless loop
	l := 0
	for {
		l++
		if l == 5 {
			break
		}
	}

	// for item of slices
	colors := []string{"blue", "white", "red", "green"}
	// run for
	for i := 0; i < len(colors); i++ {
		fmt.Println("Color:", colors[i])
	}

	//range clause
	fmt.Println("With range:")
	for i, color := range colors {
		fmt.Println(i, "-", color)
	}

	// without blank-identifier
	count := 0
	for range colors {
		count++
	}

	fmt.Println("Total elements:", count)

	newSlice := make([]int, 4)
	for i, _ := range colors {
		newSlice[i] = i
	}
	fmt.Println("New slices:", newSlice)

	// rage for string
	st := "We are using a range"
	for _, item := range st {
		fmt.Printf("%c- ", item)
	}

	// range for maps
	cars := map[string]string{"brand": "Audi", "model": "Q8", "color": "black", "location": "Germany"}

	for key, value := range cars {
		fmt.Println(key + ": " + value)
	}

	// nested loops
	fmt.Printf("Nested loops")
	parents := []int{3, 2, 1}
	child := []string{"a", "b", "c"}

	for _, i := range parents {
		fmt.Println("\n", i)
		for _, c := range child {
			fmt.Print(c + "-")
		}
	}
}
