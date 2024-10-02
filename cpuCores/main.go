package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
	"sync"
)

func main() {
	globalVariable := 1
	var wg sync.WaitGroup
	channel := make(chan string, 8)
	ch2 := make(chan int, 8)
	wg.Add(8)
	startTime := time.Now()
	var numCores = flag.Int("n", 1, "number of CPU cores to use")
	// cores := runtime.NumCPU()
	cores := 2
	fmt.Printf("This machine has %d CPU cores. \n", cores)
	runtime.GOMAXPROCS(*numCores)
	for i:=1; i<=8; i++ {
		go delayInterval(i, &wg, channel, ch2)
		fmt.Println(<-channel)
		globalVariable += <- ch2
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// asyMap := map[string]string{
	// 	"channel1": "THis is channel 1",
	// 	"channel2": "THis is channel 2",
	// }
	// go asyChannel(channel, asyMap)
	// fmt.Println(<-channel)
	wg.Wait()
	fmt.Printf("Program finished execution at %d \n", duration)
	fmt.Printf("The final value of GLOBAL VARIABLE is %d", globalVariable)
}

func delayInterval(child int, wg *sync.WaitGroup, ch chan string, ch2 chan int) {
	fmt.Printf("I am running delayed goroutine %d \n", child)
	time.Sleep(5 * time.Second)
	fmt.Printf("Completed goroutine %d \n", child)
	ch <- fmt.Sprintf("Goroutine %d is completed --- message from the channel", child)
	ch2 <- child
	wg.Done()
}

// Demonstrating the SYNCHHROUNOUS(UNBUFFERED) & ASYNCHRONOUS(BUFFERED) CHANNELS IN GO
// func asyChannel(channel chan string, channelParams map[string]string) {
// 	for _, value := range channelParams {
// 		channel <- value
// 	}
// }