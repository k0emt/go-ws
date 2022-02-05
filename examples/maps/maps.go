package main

import (
	"fmt"
	"strings"
)

func main() {

	myMap := map[string]float64{
		"FBK":     800.00,
		"Amzon":   900.00,
		"NotFlex": 777.77,
		"Gaggle":  555.44,
		"Scrub":   11.11,
	} // the trailing , is required in a multi-line definition
	// no comments within assignment

	myMap["Apl"] = 888.88
	delete(myMap, "Scrub")

	// just the keys
	for k := range myMap {
		fmt.Println(k)
	}
	fmt.Println("----")

	// just the values
	justTheValues(myMap)
	fmt.Println("----")

	// everything
	for k, v := range myMap {
		fmt.Println(k, v)
	}
	fmt.Println("----")

	value, found := myMap["DNE"] // if the given key doesn't exist, it returns 0
	fmt.Println(found)
	if !found {
		fmt.Println("No entry for given key")
	} else {
		fmt.Println(value)
	}

	myWordCounts := wordCount(`How now brown cow
	THE COW was brown
	The dog was blue
	How did the cow jump over the moon
	`)

	fmt.Println(myWordCounts)
}

func justTheValues(m map[string]float64) {
	for _, v := range m {
		fmt.Println(v)
	}
}

func wordCount(text string) map[string]int {
	words := strings.Fields(text)
	counts := map[string]int{}
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	return counts
}
