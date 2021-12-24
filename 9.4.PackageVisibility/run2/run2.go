package main

import (
	"github.com/triumphpc/go-from-simple-to-great/go-from-simple-to-great/9.4.PackageVisibility/logger2"
	"time"
)

func main() {
	logger := logger2.Instance(time.RFC3339, false)
	//logger.Log("To many info.")
	logger.Log("info", "It's info level")
	logger.Log("warning", "It's warning level")
	logger.Log("error", "It's error level")

}
