package main

import (
	"fmt"
)

// Exported type should have a comment
// Doggy is my thing that holds identifying characterists
type Doggy struct {
	Breed            string
	Weight           int
	Name             string
	Sound            string
	owner            // fieldname same as type
	secondaryContact owner
}

// d Doggy is known as the receiver
// comment is required here, because it is an exported function
// Speak is the sound that the dog makes
func (d Doggy) Speak() {
	fmt.Println(d.Sound)
}

// SpeakTwice is a double presentation of Sound
func (d Doggy) SpeakTwice() {
	d.Sound = fmt.Sprintf("%v %v", d.Sound, d.Sound) // we are only modifying our copy of d Doggy
	fmt.Println(d.Sound)
}

// using a pointer receiver here
func (d *Doggy) SetName(newName string) {
	d.Name = newName
}
