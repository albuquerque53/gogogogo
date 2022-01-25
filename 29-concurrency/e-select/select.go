package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "First Channel"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Second Channel"
		}
	}()

	//for {
	//	firstMessage := <-channel1 // This will wait the second goroutine finish to execute
	//	fmt.Println(firstMessage)

	//	secondMessage := <-channel2
	//	fmt.Println(secondMessage)
	//}

	for {
		select {
		case firstMessage := <-channel1: // Now this will run according with goroutine execution finish
			fmt.Println(firstMessage)
		case secondMessage := <-channel2:
			fmt.Println(secondMessage)
		}
	}
}
