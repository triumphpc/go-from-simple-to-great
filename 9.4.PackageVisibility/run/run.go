package main

import (
	"github.com/triumphpc/go-from-simple-to-great/go-from-simple-to-great/9.4.PackageVisibility/logger"
)

func main() {
	logger.Log("Info")

	logger.Debug(true)
	logger.Log("This is a debug")

	// Compilation error

	//fmt.Println(logger.debug)
}
