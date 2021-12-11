package main

import "fmt"

func invert(number int) int {
    return number * -1
}

func invertWithPointer(number *int) {
    *number = *number * -1
}

func main() {
    number := 10

    invertedNumber := invert(number)

    fmt.Println(invertedNumber)
    fmt.Println(number)

    anotherNumber := 30
    fmt.Printf("Before invert: %d", anotherNumber)
    invertWithPointer(&anotherNumber)
    fmt.Printf("After invert: %d", anotherNumber)
}

