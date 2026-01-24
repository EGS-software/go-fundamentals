package main

import (
	"fmt"
	"sync"
)

func main() {
	go func() {
		fmt.Printf("HI!!")
	}()

	fmt.Println("Fim da Main") // Is faster than the "go func" above

	// counters:
	var wg sync.WaitGroup // Created var wg of type sync.WaitGroup

	wg.Add(1) // Add one goroutine to the counter
	go func() {
		defer wg.Done() // Decrease the counter by one when the goroutine completes
		fmt.Printf("Hello from goroutine!")
	}()
	wg.Wait() // Wait until the counter is zero
	fmt.Printf("The End!")
}
