package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "./temp.txt"

	fileStat(fileName) // no file

	writeFile(fileName)
	readFile(fileName) // defer says wait until everything else is done

	fileStat(fileName) // file

	defer removeTempFile(fileName)
}

func fileStat(fileName string) {
	fi, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("stat failed for ", fileName)
		return
	}
	fmt.Printf("stat %v %v\n", fi.Name(), fi.Size())
}

func writeFile(fileName string) {
	content := "File writing with Go!"
	myFile, err := os.Create(fileName)
	checkError(err)

	defer myFile.Close() // this is like a finally statement
	// defer statement typically comes right after the initialization of the resource it will cleanup

	length, err := io.WriteString(myFile, content)
	checkError(err)

	fmt.Printf("Wrote file with %v characters\n", length)
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

func removeTempFile(fileName string) {
	if os.Remove(fileName) != nil {
		fmt.Println("Error removing file ", fileName)
	}
}
