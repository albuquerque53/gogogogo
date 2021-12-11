package main

import "fmt"

type user struct {
    name string
    age uint8
}

func (user user) sayHello() {
    fmt.Printf("\nHello, my name is %s", user.name)
}

func (user user) sayAge() {
    fmt.Printf("\nI'm %d years old!", user.age)
}

func (user *user) birthDay() {
    fmt.Println("\nToday is my birthday!!")
    user.age++
}

func main() {
    user := user{"Gabriel Albuquerque", 20}

    user.sayHello()
    user.sayAge()
    user.birthDay()
    user.sayAge()
}

