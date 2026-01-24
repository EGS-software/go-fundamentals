package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a new unbuffered channel that transmits strings
	messages := make(chan string)

	// Declare a WaitGroup variable to synchronize goroutines
	var wg sync.WaitGroup

	// Increment the WaitGroup counter by 1 to indicate we are waiting for one goroutine
	wg.Add(1)

	// Start an anonymous function as a goroutine (concurrent execution)
	go func() {
		// Decrement the counter immediately.
		// Note: Usually, this goes at the end (defer wg.Done()), but here it unblocks main early.
		wg.Done()

		// Send the string "Hello World!" into the channel.
		// This line will block until a receiver is ready.
		messages <- "Hello World!"
	}()

	// Block the main function until the WaitGroup counter is zero
	wg.Wait()

	// Receive the message from the channel and store it in the variable 'message'
	// This creates a synchronization point with the sender.
	message := <-messages

	// Print the received message to the console
	fmt.Println(message)

	// Print the final message
	fmt.Printf("The End!")
}
