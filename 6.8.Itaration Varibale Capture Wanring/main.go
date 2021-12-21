//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 6.8: Iteration Variable Capture Warning
//
// @author Alex Versus 2021
//

package main

import "os"

// Entrypoint
func main() {
	var delDirs []func()
	for _, d := range tempDirs() {
		dir := d // it's need!

		os.MkdirAll(dir, 0755)             // Creating
		delDirs = append(delDirs, func() { // anonymous
			os.RemoveAll(dir)
		})
	}

	// call to del dir
	for _, rmDir := range delDirs {
		rmDir()
	}

	// second example
	td := tempDirs()
	for i := 0; i < len(td); i++ {
		os.MkdirAll(td[i], 0755) // ok
		os.RemoveAll(td[i])      // warning
	}
}

func tempDirs() []string {
	return []string{
		"temp_1",
		"temp_2",
		"temp_3",
	}
}
