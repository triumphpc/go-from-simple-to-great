package go_fstg_modexam

import (
	"errors"
	"fmt"
)

// Hi example for output
func Hi(name, info string) (string, error) {
	if info == "" {
		return "", errors.New("unknown information")
	}
	return fmt.Sprintf("Hello, %s!!!, \nInformation for you:%s", name, info), nil
}