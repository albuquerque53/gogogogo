package main

import "fmt"

func main() {
	// We can call functions and assign the return to a variable
	sumNumbers := sum(10, 20)
	fmt.Println(sumNumbers)

	// Functions also are "type", so we can assign a function to a variable
	var sub = func(number1 int8, number2 int8) int8 {
		return number1 - number2
	}

	// Now we can call the variable sub of type "func"
	subNumbers := sub(30, 15)
	fmt.Println(subNumbers)

	// Multiple return needs multiple vars to assign
	calculatorSum, calculatorSub := calculator(20, 10)
	fmt.Println(calculatorSum, calculatorSub)

	// In the case of you want use just one return of multiple return function
	onlySum, _ := calculator(85, 15)
	fmt.Println(onlySum)
}

//			 argmt	 type  argmt   type  return
func sum(number1 int8, number2 int8) int8 {
	return number1 + number2
}

// A function can return multiple values:

//							arg      arg					freturn sreturn
func calculator(number1, number2 int8) (int8, int8) { // When not assign the type of arg, the type of last will be assigned
	sum := number1 + number2
	sub := number1 - number2

	return sum, sub
}
