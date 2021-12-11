package main

import "fmt"

var number int;

func init() {
    fmt.Println("Hello from init function!")
    number = 10;
}

func main() {
    fmt.Println("Hello from main function!")
    fmt.Println(number)
}

