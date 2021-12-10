package main

import "fmt"

func sum(numbers ...int) int {
    var sum int

    for _, number := range numbers {
        sum += number
    }

    return sum
}

func write(messages ...string) {
    for _, message := range messages {
        fmt.Println(message)
    }
}

func main() {
    sumOfValues := sum(1, 2, 3, 100) // return a slice

    fmt.Println(sumOfValues)
    write("Hello, World", "Привет, мир", "Olá, Mundo")
}
