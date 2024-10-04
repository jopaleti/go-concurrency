package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore struct that simulates a mutex using a buffered channel
type Semaphore struct {
	ch chan struct{} // Buffered channel to act as the semaphore
}

// NewSemaphore creates a new semaphore
func newSemaphore(limit int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, limit), // limit is the buffered size (1 for mutex)
	}
}

// Locks locks the semaphore
func (s *Semaphore) Lock() {
	s.ch <- struct{}{} // Send to channel (wait if full)
}

// Unlock unlocks the semaphore
func (s *Semaphore) Unlock() {
	<-s.ch // Release the channel
}

var sharedResource int // Shared Resource
var wg sync.WaitGroup // WaitGroup to wait for goroutines to finish

func criticalSection(sem *Semaphore, id int) {
	defer wg.Done()

	// Lock the semaphore to simulate a mutex
	sem.Lock()

	// Critical section starts
	fmt.Printf("Goroutine %d entering critical section\n", id)
	sharedResource++
	time.Sleep(1 * time.Second) // Simulate some work
	fmt.Printf("Goroutine %d leaving critical section, shared resource: %d\n", id, sharedResource)
	// Critical section ends

	// Unlock the semaphore
	sem.Unlock()
}

func main() {
	sem := newSemaphore(1) // Create a semaphore with a limit of 1 (acts as a mutex)
	
	// Start multiple goroutines that try to access the shared resource
	for i:=0; i<=5; i++ {
		wg.Add(1) // Increment waitgroup counter
		go criticalSection(sem, i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println(sharedResource)
}