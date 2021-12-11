package main

import "fmt"

func closure() func() {
    message := "Hello, World"

    function := func() {
        fmt.Println(message)
    }

    return function
}

func main() {
    message := "Olá, mundo!"
    fmt.Println(message)

    closureFunction := closure()

    closureFunction()
}
