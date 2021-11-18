package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(doSomething("aaaaAAAA"))
}

func doSomething(text string) (string, *string) {
	// error condition
	// errorMessage := "say what!"
	// return "moo", &errorMessage

	// wrong text, no error
	// return "moo", nil

	// pass the test
	return strings.ToUpper(text), nil
}
