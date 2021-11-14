package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)

	fmt.Println(states)

	// assignment
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["NY"] = "New York"
	states["CA"] = "California"
	states["NB"] = "nope"

	// deletion
	delete(states, "NB")

	fmt.Println(states)

	california := states["CA"]

	fmt.Println(california)

	// looping through k,v of states
	for abbrev, long := range states {
		fmt.Printf("%v: %v\n", abbrev, long)
	}

	// get some ordered keys
	stateKeys := make([]string, len(states))
	i := 0
	for k := range states {
		stateKeys[i] = k
		i++
	}
	sort.Strings(stateKeys)
	fmt.Println(stateKeys)

	// use the ordered keys
	for i := range stateKeys {
		fmt.Println(states[stateKeys[i]])
	}

	// an instance of Dog
	poodle := Dog{"big", 150, "Bowser"} // match order in the type
	fmt.Println(poodle)                 // only shows values
	fmt.Printf("%+v\n", poodle)         // shows field name along with value
	fmt.Printf("Breed: %v\nWeight: %v\nName: %v\n", poodle.Breed, poodle.Weight, poodle.Name)
}

// Exported type should have a comment
// needs to be exported so can use up above in main()
type Dog struct {
	Breed  string
	Weight int
	Name   string
}
