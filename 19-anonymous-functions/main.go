package main

import "fmt"

func main() {

    func() {
        fmt.Println("Hello, World!")
    }()

    hi := func() {
        fmt.Println("Olá, mundo!")
    }

    hi()

    func(name string) {
        fmt.Printf("Hello, %s", name)
    }("Albuquerque")
}
