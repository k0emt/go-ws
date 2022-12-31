package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(v int) (int, int) {
	sum := 0
	for i := 0; i < v; i++ {
		sum += v
	}
	return v, sum
}

func goRoutinesExample() {
	testValues := []int{2500, 3000, 4000}

	var waitGroup sync.WaitGroup
	for _, testValue := range testValues {
		waitGroup.Add(1)
		go func(tVal int) {
			v, sum := doSomething(tVal)
			fmt.Println(v, sum)
			waitGroup.Done()
		}(testValue)
	}
	waitGroup.Wait()
}

// channels are kind of like message queues
// consider if instead of looping through a set of numbers
// we were to have a list of URLs.
// then each go function would send the body to the channel
// where we could have go routines parsing the content
func channelExample() {
	ch := make(chan int)

	go func() {
		for i := 42; i < 45; i++ {
			fmt.Println("Sending: ", i)
			ch <- i
		}
		close(ch)
	}()

	for rcvd := range ch {
		fmt.Println("Received: ", rcvd)
	}

}

func multiChannelSelect() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 55
		ch2 <- 72 // select case short circuits, so we don't see 72
	}()

	select {
	case val := <-ch1:
		fmt.Printf("mcs got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("mcs got %d from ch2\n", val)
	}

	// part 2 -----

	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 6.28
	}()
	select {
	case val := <-out:
		fmt.Printf("mcs got %f\n", val)
	case <-time.After(200 * time.Millisecond): // change 200 to 20 to see timeout
		fmt.Println("TIMEOUT!")
	}

}

func main() {
	goRoutinesExample()
	channelExample()
	multiChannelSelect()
}
