package main

import (
	"fmt"
	"time"
)

/*
	WHAT IS SEMAPHORE IN COMPUTING
	A SEMAPHORE is a synchronisation mechanism used in computing to
	control access to shared resources by multiple processes or threads
*/

type semaphore struct {
	ch chan struct{} // Buffered channel to simulate the semaphore
}

// NewSemahore initializes a new semaphore with the given capacity
func newSemaphore(capacity int) *semaphore {
	return &semaphore{
		ch: make(chan struct{}, capacity),
	}
}

// Acquire simulates acquiring the semaphore by sending to the channel
func (s *semaphore) Acquire() {
	s.ch <- struct{}{} // Block if the channel is full
}

// Release simulates releasing the semaphore by receiving from the channel
func (s *semaphore) Release() {
	<- s.ch // Receive from the channel to release the semaphore
}

func worker(id int, sem *semaphore) {
	// simulate acquiring the semaphore
	fmt.Printf("Worker %d: Waiting to acquire semaphore...\n", id)
	sem.Acquire()
	fmt.Printf("Worker %d: Acquired semaphore, doing work...\n", id)

	// simulate work by sleeping
	time.Sleep(2*time.Second)

	// release the semaphore after work is done
	fmt.Printf("Worker %d: Releasing semaphore...\n", id)
	sem.Release()
}

func main() {
	// create a semaphore with a capacity of 2
	sem := newSemaphore(2)

	// simulate 5 workers trying to acquire the semaphore
	for i:=0; i<5; i++ {
		go worker(i, sem)
	}

	// wait for all goroutines to finish
	time.Sleep(10*time.Second)
}