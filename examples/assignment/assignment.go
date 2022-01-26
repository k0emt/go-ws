package main

import (
	"fmt"
)

const TAU = 6.28

var PI = 3.14

const (
	golden = 1.618
	kk     = 2000
)

func main() {

	a, b, c := 1, 2, 3
	var (
		d = 4
		e = 5
		f = 6
	)

	specialCharsIncluded := `She said, "don't quote me!"`

	fmt.Println(TAU, PI)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(golden, kk)
	fmt.Println(specialCharsIncluded)
}
