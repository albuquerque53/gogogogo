package main

import "fmt"

func main() {
    slice := make([]float32, 10, 12)

    // The code above will
    // 1. Create an float32 (internal) array with 12 positions
    // 2. Point 10 first indexes to slice var

    fmt.Println(slice)
    fmt.Println(len(slice))
    fmt.Println(cap(slice))

    slice = append(slice, 10.20)
    slice = append(slice, 20.40)

    // Reached the limit of 15 positions, of internal array
    fmt.Println(len(slice))
    fmt.Println(cap(slice))

    // But if we "force" one more item...
    slice = append(slice, 09.20)

    // The array positions will double
    fmt.Println(len(slice))
    fmt.Println(cap(slice))
}
