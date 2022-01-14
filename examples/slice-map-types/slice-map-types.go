package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceExample()
	mapExample()
}

func sliceExample() {
	tla := []string{"ABC", "123", "TLA", "QRX"}

	for _, short := range tla {
		fmt.Println(short)
	}

	fmt.Println("----------------")

	tla = append(tla, "BIG")

	fmt.Println(tla)

	// super easy to swap
	tla[0], tla[3] = tla[3], tla[0]

	// removing an item from a slice is gross.
	tla = removeItem(tla, "123")

	for i, short := range tla {
		fmt.Println(i, short)
	}

	fmt.Println("----------------")
}

func removeItem(sourceSlice []string, target string) []string {
	filteredSlice := []string{}

	for _, item := range sourceSlice {
		if item != target {
			filteredSlice = append(filteredSlice, item)
		}
	}

	return filteredSlice
}

func mapExample() {
	var colorMap = map[string]string{"white": "#FFFFFF", "black": "#000000", "red": "#FF0000"}
	fmt.Println(colorMap)

	// map
	// states := make(map[string]string)
	var states = map[string]string{"IN": "Indiana"}

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

	fmt.Println("----------------")

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

	fmt.Println("----------------")

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
