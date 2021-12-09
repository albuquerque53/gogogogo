package main

import "fmt"

func main() {
    number := 10

    if number > 5 {
        fmt.Println("Bigger than 5")
    } else {
        fmt.Println("Minor or equal to 5")
    }

    if otherNumber := number; otherNumber > 100 {
        fmt.Println("Bigger Than 100")
    } else if otherNumber > 10 {
        fmt.Println("Bigger Than 10")
    } else if otherNumber > 5 {
        fmt.Println("Bigger Than 5")
    } else {
        fmt.Println("Minor or equal to 5")
    }
}
