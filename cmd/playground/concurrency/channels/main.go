package main

import (
	"fmt"
	"sync"
)

func main() {
	messages := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		messages <- "Hello World!"
	}()
	wg.Wait()

	fmt.Printf("The End!")
}
