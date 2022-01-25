package main

import "fmt"

// Go Concurrency https://www.golang-book.com/books/intro/10

func main() {
	/**
	 * // First example (default way):
	 * write("Hello, World!")
	 * write("Olá, Mundo!") // This line will never run, because will wait the line above end
	 */

	// Second example (Concurrency/Goroutine):
	go write("Hello, World!") // "Hey Go, don't wait this instruction end to execute the next line"
	write("Olá, Mundo!")
}

func write(message string) {
	for { // Infinite loop
		fmt.Println(message)
	}
}
