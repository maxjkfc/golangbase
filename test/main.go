package main

import (
	"errors"
)

func main() {
	Hello(true)

}

// This is a Hello function
func Hello(flag bool) (string, error) {
	if flag {
		return "Hello", nil
	}

	return "", errors.New("No Hello")
}

