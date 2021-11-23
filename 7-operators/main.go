package main

import "fmt"

func main() {
	fmt.Println("OPERATORS:")

	// =-=-=-= Arithmetic =-=-=-=
	fmt.Println("=-=-=-= Arithmetic =-=-=-=")
	sum := 1 + 2
	sub := 1 - 2
	div := 2 / 1
	mul := 2 * 2
	mod := 3 % 2
	fmt.Println(sum, sub, div, mul, mod)
	var number1 int16 = 20
	var number2 int32 = 100
	fmt.Println(number1, number2)
	// number3 := number1 + number2 (error mismatched types)

	// =-=-=-= Allocation =-=-=-=
	fmt.Println("=-=-=-= Allocation =-=-=-=")
	var typeAllocation string = "Hi!"
	simpleAllocation := "Hello!"
	fmt.Println(typeAllocation, simpleAllocation)

	// =-=-=-= Relation (bools) =-=-=-=
	fmt.Println("=-=-=-= Relation =-=-=-=")
	fmt.Println(2 > 1)
	fmt.Println(2 >= 1)
	fmt.Println(2 == 1)
	fmt.Println(2 < 1)
	fmt.Println(2 <= 1)
	fmt.Println(2 != 1)

	// =-=-=-= Logical =-=-=-=
	fmt.Println("=-=-=-= Logical =-=-=-=")
	truthy, falsy := true, false
	fmt.Println(truthy && falsy) // AND
	fmt.Println(truthy || falsy) // OR
	fmt.Println(!truthy)         // NOT

	// =-=-=-= Unary =-=-=-=
	fmt.Println("=-=-=-= Unary =-=-=-=")
	var one int8 = 1
	one++
	fmt.Println(one)
	one += 10
	fmt.Println(one)
	one--
	fmt.Println(one)
	one -= 10
	fmt.Println(one)
	one *= 10
	one /= 2
	one %= 2
	fmt.Println(one)
}
