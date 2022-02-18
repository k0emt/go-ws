package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n==============================================\n", resp)

	// do it all myself and put ZZZ after each buffer chunk
	/*
		buff := make([]byte, 498)
		byteCount, _ := resp.Body.Read(buff)
		for byteCount > 0 {
			fmt.Printf("%v ZZZ\n", string(buff))
			byteCount, _ = resp.Body.Read(buff)
		}
	*/

	// io.Copy(os.Stdout, resp.Body)

	cw := customWriter{}
	io.Copy(cw, resp.Body)
}

func (customWriter) Write(buff []byte) (int, error) {
	fmt.Printf("%d | %v\n\n", len(buff), string(buff))

	return len(buff), nil
}
