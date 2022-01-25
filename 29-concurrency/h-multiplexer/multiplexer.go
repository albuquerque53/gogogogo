package main

import "fmt"

// Multiplexer will merge two channels and provide an outputChannel

func main() {
	outputChannel := mergeChannels(write("Hello, World"), write("Hello, Go"))

	for counter := 0; counter < 10; counter++ {
		fmt.Println(<-outputChannel)
	}
}

func mergeChannels(firstEntryChannel, secondEntryChannel <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-firstEntryChannel:
				outputChannel <- message
			case message := <-secondEntryChannel:
				outputChannel <- message

			}
		}
	}()

	return outputChannel
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
