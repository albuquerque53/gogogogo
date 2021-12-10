package main

import "fmt"

func firstFunction() {
    fmt.Println("First function!")
}

func secondFunction() {
    fmt.Println("Second function!")
}

func checkIfStudentIsApproved(firstTestNote, secondTestNote int) bool {
    defer fmt.Println("Finished computing! The student is:")
    fmt.Println("Computing student notes...")

    mean := firstTestNote + secondTestNote / 2

    return mean >= 6
}


func main() {
    firstFunction()
    secondFunction()
    checkIfStudentIsApproved(5, 3)
}
