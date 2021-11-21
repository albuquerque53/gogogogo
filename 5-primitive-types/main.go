package main

import (
	"errors"
	"fmt"
)

func main() {
	// =-=-=-=-=-=-=-= INTEGER =-=-=-=-=-=-=-=
	/*
			int8 (byte): -128 to 127
			int16: -32768 to 32767
			int32 (rune): -2147483648 to 2147483647
			int64: -9223372036854775808 to 9223372036854775807

		Type Inference:
			When not specified the type, go will take `int` type based on system arch (32/64)
	*/

	var number int8 = 100
	// var number int16 = 40000
	fmt.Println(number)

	/**
	Unsigned Integers:
		uint16:	0 to 65535
		uint32: 0 to 4294967295
		uint64: 0 to 2147483648
	*/

	var uintNumber uint16 = 30000
	// var uintNumber uint16 = -1
	fmt.Println(uintNumber)

	// The "zero value" of any integer is 0
	var zeroInteger int
	fmt.Println(zeroInteger) // 0

	// =-=-=-=-=-=-=-= FLOAT =-=-=-=-=-=-=-=

	var floatNumber float32 = 10.000
	fmt.Println(floatNumber)

	var secondFloatNumber float64 = 10.000000000000
	fmt.Println(secondFloatNumber)

	// The "zero value" of float is 0.0
	var zeroFloat float32
	fmt.Println(zeroFloat) // 0.0

	// =-=-=-=-=-=-=-= STRING =-=-=-=-=-=-=-=

	var myString string = "Text"
	fmt.Println(myString)

	// Always use " for strings and not ' (simple quotes means "char")
	char := '%'
	fmt.Println(char) // 37 - ASCII of %

	// The "zero value" of string is ""
	var zeroString string
	fmt.Println(zeroString) // "" (empty string)

	// =-=-=-=-=-=-=-= BOOLEAN =-=-=-=-=-=-=-=

	var success bool = true
	fmt.Println(success)

	var fail bool = false
	fmt.Println(fail)

	// The "zero value" of boolean in false
	var zeroBoolean bool
	fmt.Println(zeroBoolean) // false

	// =-=-=-=-=-=-=-= ERROR =-=-=-=-=-=-=-=

	var exception error = errors.New("Internal Error")
	fmt.Println(exception)

	// The "zero value" of error in <nil>
	var zeroError error
	fmt.Println(zeroError) // <nil>

}
