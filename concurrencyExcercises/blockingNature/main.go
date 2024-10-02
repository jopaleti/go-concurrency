package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Creating a new channel
	ch := make(chan string, 3)
	wg.Add(1)
	// Starting a goroutine that receives a channel after 5 seconds 
	go func () {
		defer wg.Done()
		for i:=0; i<3; i++ {
			fmt.Println("Goroutine: Waiting for 5 seconds before receiving from the channel...")
			time.Sleep(5*time.Second) // SImulate waiting for 5 seconds
			msg := <- ch
			fmt.Println("Goroutine: Received from channel: ", msg)
		}
	}()

	msg := [3]string{"sending", "received", "sent"}
	for i:=0; i<3; i++ {
		// Put a value on the channel (this will block until the goroutine receives it)
		fmt.Println("Main: Sending value to the channel...")
		ch <- msg[i]
		fmt.Println("Main: Value sent to the channel.")
		// time.Sleep(7*time.Second)
	}
	// Allow time for the goroutine to finish before main exits
	wg.Wait()
}