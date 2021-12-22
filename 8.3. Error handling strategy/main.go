//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 8.3: Error handling
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Entrypoint
func main() {
	//t := time.Date(2021, 1, 1, 1, 1, 1, 1, nil) // error
	//fmt.Println(t)

	// Create new error
	err := errors.New("new error")
	fmt.Println("Print new error:", err)
	//
	// Dynamic error string
	err = fmt.Errorf("error at: %v", time.Now())
	fmt.Println("An error handling:", err)

	err = createError()
	if err != nil {
		fmt.Println("Again", err)
		return
	}
	//
	// first strategy example
	_, err = callFirst("test")
	if err != nil {
		fmt.Println("First strategy error", err)
	}
	//
	//// first strategy example
	//_, err = callSecond("test")
	//if err != nil {
	//	fmt.Println("First strategy error", err)
	//}
	//
	// second strategy example
	//err = callSecondStrategy("test")
	//if err != nil {
	//	fmt.Println("Second strategy error", err)
	//}

	//// third strategy example
	//err := callThirdStrategy("test")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	//	os.Exit(1)
	//}

	// forth strategy
	//err := callThirdStrategy("test")
	//if err != nil {
	//	log.Fatalf("Server error %v\n", err)
	//}

	// fifth strategy
	//err := callThirdStrategy("test")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	//}

	os.RemoveAll("/tmp_errors") // ignore RemoveAll erros

}

// Create new error
func createError() error {
	return errors.New("error")
}

// Example several return values
func example(n string) (string, error) {
	if n != "" {
		return strings.Title(n), nil
	}
	return "", errors.New("some error")
}

// Failure of the calling function
func callFirst(url string) (string, error) {
	_, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return "some answer", nil
}

// Failure of the calling function
func callSecond(url string) (string, error) {
	_, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("additional %s info. Error: %v", url, err)
	}
	return "some answer", nil
}

func callSecondStrategy(url string) error {
	const timeout = 5 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("url %s is not available...%v", url, err)
		time.Sleep(time.Second << uint(tries)) // Increase
	}
	return fmt.Errorf("Server %s is not response; time %s ", url, timeout)
}

func callThirdStrategy(url string) error {
	const timeout = 5 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("url %s is not available...%v", url, err)
		time.Sleep(time.Second << uint(tries)) // Increase
	}
	return errors.New("after 5 second process stopped")
}
