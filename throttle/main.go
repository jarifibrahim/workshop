package main

import (
	"fmt"
	"runtime"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func doWork(th *Throttle) {
	// Mark this job as done.
	defer th.Done(nil)

	// Simulate some work.
	time.Sleep(time.Second)
}

func printStats() {
	for {
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Active number of goroutines = %+v\n", runtime.NumGoroutine())
	}
}

func main() {
	go printStats()

	// Create a new throttle variable.
	th := NewThrottle(3)

	for i := 0; i < 100; i++ {
		// Gate keeper. Do not let more than 3 goroutines to start.
		th.Do()

		fmt.Printf("Processing item number = %+v\n", i)
		go doWork(th)
	}

	// Wait for all the jobs to finish.
	th.Finish()
}
