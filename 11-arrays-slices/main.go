package main

import "fmt"

func main() {
    fmt.Println("=-=-=-= Array =-=-=-=")
    // Simple declaration
    var exampleArray[3] string
    fmt.Println(exampleArray)

    exampleArray[0] = "Hello"
    exampleArray[1] = "Oi"
    exampleArray[2] = "Привет"

    fmt.Println(exampleArray)

    // Type inference declaration
    otherArray := [2]bool{true, false}
    fmt.Println(otherArray)

    // Error (invalid index)
    // otherArray[2] = false

    // Without specify the array size
    sizeArray := [...]int{1, 2 , 3}
    fmt.Println(sizeArray)

    // Error (invalid index)
    // sizeArray[3] = 4

    fmt.Println("=-=-=-= Slice =-=-=-=")
    mySlice := []int{201, 202, 203}
    fmt.Println(mySlice)

    mySlice = append(mySlice, 100)
    // More flexible
    fmt.Println(mySlice)

    // Pointing otherSlice to index 1 to 2 of exampleArray
    otherSlice := exampleArray[1:3]
    fmt.Println(otherSlice)

    // If changes the original
    exampleArray[1] = "Ciao"
    fmt.Println(otherSlice) // Affects the pointer
}
