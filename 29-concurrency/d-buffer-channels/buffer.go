package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Hello, World!"

	message := <-channel

	fmt.Println(message)
}
