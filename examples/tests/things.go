package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(doSomething("aaaaAAAA"))
}

// returning upper cased text, error message on empty string
func doSomething(text string) (string, *string) {

	if len(text) == 0 {
		expected := "need to provide text to operate on"
		return "", &expected
	}

	return strings.ToUpper(text), nil
}

// returning lower cased text, boolean flag to indicate success in conversion
func doAthing(text string) (string, bool) {
	return strings.ToLower(text), len(text) != 0
}
