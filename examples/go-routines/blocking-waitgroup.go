package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sleepers := []int{1, 2, 3, 2, 2, 3, 4, 3, 1}
	var wg sync.WaitGroup // use this construct to make sure everything has completed

	start := time.Now()

	fmt.Println("Starting at:", start)

	for _, sleepTime := range sleepers {
		wg.Add(1)
		go wasteTime(&wg, sleepTime)
	}
	fmt.Println("Main: WAITING for go routines to complete")
	wg.Wait()

	end := time.Now()
	fmt.Println("Main: all go routines COMPLETED")
	fmt.Println("Completed at:", end)
	fmt.Println("Elapsed: ", time.Since(start))
}

func wasteTime(wg *sync.WaitGroup, seconds int) {
	defer wg.Done()

	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Println("Completed:", seconds)
}
