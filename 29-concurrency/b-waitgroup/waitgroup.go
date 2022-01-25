package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		writeOnScreen("Hello, World!")
		waitGroup.Done() // -1
	}()

	go func() {
		writeOnScreen("Olá, Mundo!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait()
}

func writeOnScreen(message string) {
	for counter := 0; counter < 5; counter++ {
		fmt.Println(message)
	}
}
