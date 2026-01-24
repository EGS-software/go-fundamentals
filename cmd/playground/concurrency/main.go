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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Hello from goroutine!")
	}()
	wg.Wait()
	fmt.Printf("The End!")
}
