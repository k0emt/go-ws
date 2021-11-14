package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "./temp.txt"

	writeFile(fileName)
	defer readFile(fileName) // defer says wait until everything else is done
}

func writeFile(fileName string) {
	content := "File writing with Go!"
	myFile, err := os.Create(fileName)
	checkError(err)

	length, err := io.WriteString(myFile, content)
	checkError(err)

	fmt.Printf("Wrote file with %v characters\n", length)

	defer myFile.Close()
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file: ", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
