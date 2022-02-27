package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// take a file name to read as arg 1
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("usage is %v filename\n", os.Args[0])
		return
	}

	fileName := os.Args[1]

	// the easy way ----------------------------------------------
	// data, _ := ioutil.ReadFile(fileName)
	// fmt.Printf("Text read from file:\n\n%v\n\n", string(data))

	// now at a lower level --------------------------------------

	// open the file
	// file, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
	file, err := os.Open(fileName) // in effect, same as above

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	cw := customWriter{counter: 0}
	io.Copy(&cw, file)
}

type customWriter struct {
	counter int
}

func (cw *customWriter) Write(buff []byte) (int, error) {
	// it's a simple matter to just write out string(buff) here
	// let's do something more fun

	lines := strings.Split(string(buff), "\n")
	for _, line := range lines {
		cw.counter++
		fmt.Printf("%5d | %3d [%v] %v\n", cw.counter, len(line), line, getBiggestWord(line))
	}

	return len(buff), nil
}

func getBiggestWord(line string) string {
	words := strings.Split(line, " ")
	biggestWord := ""
	for _, word := range words {
		thisWord := strings.Trim(word, " \t\n.,?")
		if len(thisWord) > len(biggestWord) {
			biggestWord = thisWord
		}
	}
	return biggestWord
}
