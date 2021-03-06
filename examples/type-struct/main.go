package main

import (
	"fmt"
)

func main() {
	orderBasedAssignment()
	fmt.Println("------------")
	anotherWay()
	fmt.Println("------------")
	oneMoreWay()
}

func orderBasedAssignment() {
	// an instance of Dog
	peke := Doggy{"Pekinese", 15, "Bouncer", "grrrRr!",
		owner{"Mom", "4175551234"}, owner{}} // match order in the type
	fmt.Println(peke)         // only shows values
	fmt.Printf("%+v\n", peke) // shows field name along with value
	fmt.Printf("Breed: %v\nWeight: %v\nName: %v\n", peke.Breed, peke.Weight, peke.Name)

	peke.Speak()
	peke.SpeakTwice()
	peke.Speak()

	peke.SetName("Old Dog")
	fmt.Println(peke.Name)
}

func anotherWay() {
	yorky := Doggy{Name: "T.Rex", Weight: 6, owner: owner{name: "Dad", phone: "+15735551212"}} // named arguments and embedded
	fmt.Printf("Name: %v, weight: %v, owner: %v\n", yorky.Name, yorky.Weight, yorky.owner)
	fmt.Println(yorky)

}

func oneMoreWay() {
	var sable Doggy // default zero values empty string, and 0
	fmt.Println(sable)

	sable.Name = "Sable"
	sable.Weight = 7

	fmt.Println(sable)
}
