package main

import (
	"github.com/triumphpc/logger"
	"time"
)

func main() {
	log := logger.Instance(time.RFC3339, true)
	log.Log("info", "It's info level")
	log.Log("warning", "It's warning level")
	log.Log("error", "It's error level")
}
