package main

import "fmt"

func repeat(times int, toRepeat func()) {
	for i := 0; i < times; i++ {
		toRepeat()
	}
}

func hello() {
	fmt.Println("Hello World!")
}

func main() {

	hello()

	repeat(3, hello)

	repeat(3, func() {
		fmt.Println("anonymous function")
	})

}
