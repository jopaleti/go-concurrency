package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(8)
	startTime := time.Now()
	var numCores = flag.Int("n", 1, "number of CPU cores to use")
	// cores := runtime.NumCPU()
	cores := 2
	fmt.Printf("This machine has %d CPU cores. \n", cores)
	runtime.GOMAXPROCS(*numCores)
	for i:=1; i<=8; i++ {
		go delayInterval(i, &wg)
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	wg.Wait()
	fmt.Printf("Program finished execution at %d \n", duration)
}

func delayInterval(child int, wg *sync.WaitGroup) {
	fmt.Printf("I am running delayed goroutine %d \n", child)
	time.Sleep(5 * time.Second)
	fmt.Printf("Completed goroutine %d \n", child)
	wg.Done()
}

// Demonstrating the SYNCHHROUNOUS(UNBUFFERED) & ASYNCHRONOUS(BUFFERED) CHANNELS IN GO
