package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var result string

	myRandom := rand.Intn(5) + 1
	switch myRandom {
	case 1:
		result = "uno" // break is implicit now
	case 2:
		result = "dos"
	case 3:
		result = "tres"
		fallthrough // is like old C behavior without a break
	default:
		result += "some other number"
	}

	fmt.Println("My random: ", myRandom, " result ", result)

	switch myRand := rand.Intn(4) + 1; myRand { // number 1-4, can init within switch, sets scope to switch
	case 1:
		result = fmt.Sprintf("%d %s", myRand, "one") // myRand isn't accessible outside of the switch
	case 2:
		result = fmt.Sprintf("%d %s", myRand, "two")
	case 3:
		result = fmt.Sprintf("%d %s", myRand, "three")
	default:
		result = fmt.Sprintf("%d %s", myRand, "some other number")
	}

	fmt.Println("My random from myRand: ", result)
}
