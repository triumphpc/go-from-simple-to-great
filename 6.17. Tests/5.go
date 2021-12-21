//
// Go (Golang) From simple to great. The Complete Developer's Guide.
//
// Write a function to traverse directories recursively and list files
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	err := scanDir(".")
	if err != nil {
		log.Fatal(err)
	}
}

// Scan directory
func scanDir(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDir(path)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(path)
		}
	}
	return nil
}
