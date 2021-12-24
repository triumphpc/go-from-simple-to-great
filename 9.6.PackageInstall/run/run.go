package main

import (
	"github.com/triumphpc/go-from-simple-to-great/go-from-simple-to-great/9.6.PackageInstall/run/logger"
	"time"
)

func main() {
	log := logger.Instance(time.RFC3339, true)
	//logger.Log("To many info.")
	log.Log("info", "It's info level")
	log.Log("warning", "It's warning level")
	log.Log("error", "It's error level")
}
