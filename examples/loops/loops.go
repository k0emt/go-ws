package main

import (
	"fmt"
)

func main() {

	colors := []string{"Red", "White", "Blue", "Green"}

	fmt.Println(colors)

	for i := 0; i < len(colors); i++ { // like old school C w/out ()
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for i, color := range colors { // indexVariable, value name
		fmt.Printf("%d %s\n", i, color)
	}

	for _, color := range colors { // unused indexVariable, value name
		fmt.Println(color)
	}

	// NO while loop, variant of for
	value := 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value++
	}

	value = 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value += 2
		if value > 4 {
			goto getOuttaLoop // WOW! a goto
		}
	}

getOuttaLoop:
	fmt.Println("out of the loop")

}
