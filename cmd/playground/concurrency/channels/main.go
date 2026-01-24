package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	messages <- "Hello World!"

	fmt.Printf("The End!")
}
