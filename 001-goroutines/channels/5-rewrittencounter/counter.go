package main

import (
	"fmt"
	"sync"
)

// Function to increment the counter
func incrementWorker(times int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure WaitGroup is marked as done
	for i := 0; i < times; i++ {
		ch <- 1 // Send +1 to the channel
	}
}

// Function to decrement the counter
func decrementWorker(times int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure WaitGroup is marked as done
	for i := 0; i < times; i++ {
		ch <- -1 // Send -1 to the channel
	}
}

// Function to process the channel and update the counter
func updateCounter(ch chan int, counter *int, done chan bool) {
	for delta := range ch {
		*counter += delta // Update the counter based on the value received
	}
	done <- true // Signal that the update is complete
}

func main() {
	counter := 0            // Shared counter
	wg := sync.WaitGroup{}  // WaitGroup for synchronization
	ch := make(chan int)    // Channel for managing counter updates
	done := make(chan bool) // Channel for signaling the completion of updates

	// Start the counter updater as a separate goroutine
	go updateCounter(ch, &counter, done)

	// Start increment and decrement workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incrementWorker(1000, ch, &wg) // Increment 1000 times

		wg.Add(1)
		go decrementWorker(500, ch, &wg) // Decrement 500 times
	}

	wg.Wait() // Wait for all workers to finish
	close(ch) // Close the channel after workers are done

	<-done // Wait for the counter updater to finish
	fmt.Println("Final counter value:", counter)
}
