package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	response, err := http.Get(url)
	checkForError(err)

	fmt.Printf("Response Type: %T\n", response)

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	checkForError(err)

	content := string(bytes)
	// fmt.Print(content) // now need to parse this text into something we can work with

	tours := toursFromJson(content)

	for _, tour := range tours {
		fmt.Printf("%v,%v,%v\n", tour.Name, tour.Price, tour.Difficulty)
	}
}

// if we had a bunch of code together in the same directory
// all of the code files would see this function
func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

// toursFromJson converts JSON to a Tour slice
// TODO: candidate for a unit test
func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkForError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour) // <-- here is our Tour struct
		checkForError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price, Difficulty string
	// TODO: how to turn the Price into an integer?
}
