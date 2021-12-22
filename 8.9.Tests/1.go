/**
Go (Golang) From simple to great. The Complete Developer's Guide.

Correct the code below so that closing the file will work in any case, even if there is an error while processing the file


func main() {
	// Get data from file
	data, err := GetData("text.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data:", data)
}

func GetData(fileName string) ([]float64, error) {
	var n []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		n = append(n, data)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	defer CloseFile(file)

	return n, nil
}

// Open file function
func OpenFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

// Close file function
func CloseFile(file *os.File) {
	file.Close()
}



@author Alex Versus 2021
www.alexversus.com
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Get data from file
	data, err := GetData("text.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data:", data)
}

func GetData(fileName string) ([]float64, error) {
	var n []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	defer CloseFile(file) // in this point before main logic

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		n = append(n, data)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return n, nil
}

// Open file function
func OpenFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

// Close file function
func CloseFile(file *os.File) {
	file.Close()
}
