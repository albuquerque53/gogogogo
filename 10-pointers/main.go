package main

import "fmt"

func main() {
	// Default vars
	var firstNumber int = 10
	var secondNumber int = firstNumber

	fmt.Println(firstNumber, secondNumber)

	firstNumber++
	fmt.Println(firstNumber, secondNumber)

	// Pointer
	var thirdNumber int
	var pointer *int

	thirdNumber = 100
	// Wrong attrb
	//pointer = thirdNumber

	// Right way
	pointer = &thirdNumber

	fmt.Println(thirdNumber, pointer) // Pointers always refers to mem addr
    fmt.Println(*pointer) // output: 100 (Pointer dereference)
}
