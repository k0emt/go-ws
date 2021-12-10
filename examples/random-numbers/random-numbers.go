package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// random numbers
	rSource := rand.NewSource(time.Now().UnixNano())
	rGenerator := rand.New(rSource)

	fmt.Printf("%2d %f\n", rGenerator.Intn(10), rGenerator.Float64()*5.0) // random numbers

	fmt.Println(rGenerator.Perm(10)) // a slice of random order ints [0,n)

}
