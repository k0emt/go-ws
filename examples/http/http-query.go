package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n==============================================\n", resp)

	buff := make([]byte, 498)
	byteCount, _ := resp.Body.Read(buff)
	for byteCount > 0 {
		fmt.Printf("%v ZZZ\n", string(buff))
		byteCount, _ = resp.Body.Read(buff)
	}

	// io.Copy(os.Stdout, resp.Body)
}
