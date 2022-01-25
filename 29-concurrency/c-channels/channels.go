package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go writeOnScreen("Hello, World!", channel)

	//for {
	//	message, open := <-channel
	//	if !open { // if to prevent deadlocks
	//		break
	//	}

	//	fmt.Println(message)
	//}

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("That's all folks!")
}

func writeOnScreen(message string, channel chan string) {
	for counter := 0; counter < 5; counter++ {
		channel <- message
		time.Sleep(time.Second)
	}

	close(channel)
}
