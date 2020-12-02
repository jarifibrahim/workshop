package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doWork(cl *Closer) {
	defer cl.Done()

	i := 0
	for {
		select {
		case <-cl.HasBeenClosed():
			// Close active connection.
			// Close open file descriptors.
			fmt.Println("Exiting doWork. Bye Bye!")
			return
		default:
			// Simulate some work.
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("Processed Items = %+v\n", i)
			i++
		}
	}
}

func main() {
	cl := NewCloser(1)

	go doWork(cl)

	// Gracefully exit on CTRL+C.
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		<-sigCh
		// Signal the goroutine to stop.
		cl.Signal()
	}()

	// Wait for the goroutines to finish.
	cl.Wait()
}
