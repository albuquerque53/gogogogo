package main

import "fmt"


func genericFunction(genericInterface interface{}) {
    fmt.Println(genericInterface)
}

func main() {
    genericFunction("Anything")
    genericFunction(100)
    genericFunction(true)

    // Example of generic iface usage is fmt.Println()
    fmt.Println(1, true, "anything")

    // Never do this:
    myMap := map[interface{}]interface{} {
        "value1": true,
        12: float32(100),
        false: "myString",
    }

    fmt.Println(myMap)
}

