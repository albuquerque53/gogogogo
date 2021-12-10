package main

import "fmt"

func tryToRecover(mean float64) bool {
    fmt.Println("Trying to recover excecution...")

    if recover := recover(); recover != nil {
        fmt.Println("Execution recovered!")
        return mean >= 6
    }

    return false
}

func getMeanOfStudent(firstNote, secondNote float64) bool {
    fmt.Println("Counting the student notes...")

    mean := firstNote + secondNote / 2
    defer tryToRecover(mean)

    if mean > 6 {
        return true
    }

    if mean < 6 {
        return false
    }

    panic("THE MEAN IS EXACTLY 6")
}

func main() {
    approved := getMeanOfStudent(6, 6)

    fmt.Println(approved)
}
