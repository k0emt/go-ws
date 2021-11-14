package main

import (
	"fmt"
)

func main() {
	getToIt()
	fmt.Println("Doing maths: ", doMaths(4, 2))
	fmt.Println("Doing maths same: ", doMathsSame(4, 2))
	fmt.Println("maths all of them: ", doMathsOnAll(4, 2, 0))

	multiMath, multiCount := doMathsOnAllWithCount(4, 2, 0, 0)
	fmt.Printf("Math: %d values: %d\n", multiMath, multiCount)
}

func getToIt() {
	fmt.Println("get to it!")
	// no return value
}

func doMaths(a int, b int) int {
	return a + b
}

func doMathsSame(a, b int) int {
	return a + b
}

func doMathsOnAll(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// multiple return values, I love it!
func doMathsOnAllWithCount(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
