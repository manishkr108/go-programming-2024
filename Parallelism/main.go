package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("Letter: %c\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // Use 2 CPU cores

	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers(&wg) // run on one CPU core
	go printLetters(&wg) // run on another CPU core

	wg.Wait() // Wait for both goroutines to finish
}
