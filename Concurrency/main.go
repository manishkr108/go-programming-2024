package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(1 * time.Second) // Simulate a delay
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("Letter: %c\n", i)
		time.Sleep(1 * time.Second) // Simulate a delay
	}
}

func main() {
	go printNumbers() // run as a goroutine (concurrent)
	go printLetters() // run as a goroutine (concurrent)

	time.Sleep(6 * time.Second) // Wait for goroutines to finish
}
