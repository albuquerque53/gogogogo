package main

import "fmt"

func main() {
	// Specifying the var type on declr (explicit type)
	var name string = "Albuquerque"

	// Without specifying the var type on declr (type inference)
	role := "Developer"

	fmt.Println(name, role)

	// Multiple var declaration with explicit type
	var (
		email string = "gabrielalbuquerque-dev@hotmail.com"
		age   int    = 19
	)

	fmt.Println(email, age)

	// Multiple var declaration with type inference
	hobbie, favoriteMusic := "Play guitar", "Folsom Prison Blues"

	fmt.Println(hobbie, favoriteMusic)

	// Declaring constants
	const color string = "Gray"

	fmt.Println(color)

	// Var inversion
	hobbie, favoriteMusic = favoriteMusic, hobbie

	fmt.Println(hobbie, favoriteMusic)
}
