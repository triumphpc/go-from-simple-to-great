package logger

import (
	"fmt"
	"time"
)

var debug bool

// Exporting function
func Debug(b bool) {
	debug = b
}

// Exporting function
func Log(st string) {
	if !debug {
		return
	}
	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), st)
}
