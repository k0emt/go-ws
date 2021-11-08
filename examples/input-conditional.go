package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// get some console input and shove it into userProvidedNumber
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an integer number: ")
	input, _ := reader.ReadString('\n')
	fmt.Print("you entered: ", input)

	userProvidedNumber, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("That wasn't good Integer!")
	}

	var result string

	if userProvidedNumber < 0 { // look ma, no parens!
		result = "Less than zero"
	} else if userProvidedNumber == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)
}
