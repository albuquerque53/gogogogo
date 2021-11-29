package main

import "fmt"

type book struct {
	title  string
	author string
}

type music struct {
	name   string
	artist artist
}

type artist struct {
	name string
}

func main() {
	var myBook book

	myBook.title = "The Lord of Rings"
	myBook.author = "J.R.R Tolkien"

	fmt.Println(myBook)
	fmt.Println(myBook.title)
	fmt.Println(myBook.author)

	otherBook := book{"1984", "George Orwell"}

	fmt.Println(otherBook)

	bookWithoutName := book{author: "Sun Tzu"}

	fmt.Println(bookWithoutName)

	cash := artist{"Johnny Cash"}

	cocaineBlues := music{"Cocaine Blues", cash}

	fmt.Println(cocaineBlues.artist.name)
}
