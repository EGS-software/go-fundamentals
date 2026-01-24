package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Printf("HI!!")
	}()

	fmt.Println("Fim da Main") // Is faster than the "go func" above
}
