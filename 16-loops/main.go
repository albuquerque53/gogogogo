package main

import (
	"fmt"
)

func main() {
	//counter := 0
	//for counter < 10 {
	//	counter++
	//	fmt.Println(counter)
	//	time.Sleep(time.Second)
	//}

    //for secondCounter := 0; secondCounter < 10; secondCounter++ {
    //    fmt.Println(secondCounter)
    //}

    roles := [3]string{"Developer", "Product Owner", "Scrum Master"}

    for index, value := range roles {
        fmt.Println(index, value)
    }

    for index, letter := range "STRING" {
        // fmt.Println(index, letter) letter will be ascii referent
        fmt.Println(index, string(letter))
    }

    user := map[string]string {
        "name": "Nikola Tesla",
        "role": "Inventor, Eletrical Engineer, Scientist and Futurist",
    }

    for key, value := range user {
        fmt.Println(key, value)
    }
}
