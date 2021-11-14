package main

import (
	"fmt"
)

func main() {
	// an instance of Dog
	peke := Doggy{"Pekinese", 15, "Bouncer", "grrrRr!"} // match order in the type
	fmt.Println(peke)                                   // only shows values
	fmt.Printf("%+v\n", peke)                           // shows field name along with value
	fmt.Printf("Breed: %v\nWeight: %v\nName: %v\n", peke.Breed, peke.Weight, peke.Name)

	peke.Speak()
	peke.SpeakTwice()
	peke.Speak()
}
