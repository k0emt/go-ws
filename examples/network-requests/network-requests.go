package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	response, err := http.Get(url)
	checkForError(err)
	defer response.Body.Close()

	fmt.Printf("Response Type: %T\n", response)

	bytes, err := ioutil.ReadAll(response.Body)
	checkForError(err)

	content := string(bytes)
	// fmt.Print(content) // now need to parse this text into something we can work with

	tours := toursFromJson(content) // then maybe another step to get into struct with types we want

	for _, tour := range tours {
		fmt.Printf("%v, %v, %.2f, %v\n", tour.Title, tour.Name, tour.PriceNumeric(), tour.Difficulty)
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
	Name, Difficulty string
	Price            string // is PriceNumeric() the correct way to translate to numeric?
	Title            string `json:"packageTitle"` // map to a different field name in the json
}

func (tour *Tour) PriceNumeric() float64 {
	parsedPrice, _ := strconv.ParseFloat(tour.Price, 64)
	return parsedPrice
}
