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

func doWork() {
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

	for i := 0; i < 100; i++ {
		fmt.Printf("Processing item number = %+v\n", i)
		go doWork()
	}

	time.Sleep(time.Second * 2)
}
