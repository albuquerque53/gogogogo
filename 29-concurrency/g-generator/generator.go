package main

import "fmt"

/**
 * GENERATOR
 *
 * A function that encapsules the call for a goroutine
 * and returns a channel
 */

func main() {
	writeChannel := write("Hello, World")

	for counter := 0; counter < 10; counter++ {
		fmt.Println(<-writeChannel)
	}
}

func write(text string) <-chan string {
	myChannel := make(chan string)

	go func() {
		for {
			myChannel <- fmt.Sprintf(text)
		}
	}()

	return myChannel
}
