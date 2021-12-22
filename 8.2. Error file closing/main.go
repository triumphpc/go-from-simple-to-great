//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Example 8.2: Error file closing
//
// @author Alex Versus 2021
// www.alexversusu.com

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Entrypoint
func main() {
	numbers, err := GetFloatsFromFile("file.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Printf("Sum: %0.2f\n", sum)

}

// Open file function for writing
func openFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)

	return os.Open(fileName)
}

// Close file
func closeFile(file *os.File) {
	fmt.Println("Closing file")

	file.Close()
}

// Get content from file
func GetFloatsFromFile(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := openFile(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	// Closing file
	closeFile(file)

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}
