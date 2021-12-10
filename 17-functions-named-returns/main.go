package main

import "fmt"

func calc(number1, number2 int) (sum int, sub int) {
    sum = number1 + number2
    sub = number1 - number2

    return
}

func main() {
    sum, sub := calc(20, 10)

    fmt.Println(sum, sub)
}
