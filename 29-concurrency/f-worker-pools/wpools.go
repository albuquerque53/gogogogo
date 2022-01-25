package main

import "fmt"

func main() {
	// This will store the length of the sequence that we want to print
	length := make(chan uint, 45)
	// This will store the fibonacci sequence
	fibonacciResults := make(chan uint, 45)

	// Using 4 threads running concurrently to calc fibonacci
	go worker(length, fibonacciResults)
	go worker(length, fibonacciResults)
	go worker(length, fibonacciResults)
	go worker(length, fibonacciResults)

	// Adding values to length
	for counter := uint(0); counter < 45; counter++ {
		length <- counter
	}
	close(length)

	// Running throught results and printing
	for counter := 0; counter < 45; counter++ {
		result := <-fibonacciResults
		fmt.Printf("%d:  %d\n", counter, result)
	}
}

//          length is an exit channel    results is an input channel
func worker(length <-chan uint, fibonacciResults chan<- uint) {
	for number := range length {
		fibonacciResults <- calcFibonacci(number)
	}
}

func calcFibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return calcFibonacci(position-2) + calcFibonacci(position-1)
}
