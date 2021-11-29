package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
}

type student struct {
	person
	course string
	campus string
}

func main() {
	zeca := person{"Zeca", "Baleiro", 55}
	csStudent := student{zeca, "Computer Science", "Massachusetts"}

	fmt.Println(zeca.age)
	fmt.Println(csStudent.course)
	fmt.Println(csStudent.name)
}
