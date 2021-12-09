package main

import "fmt"

func main() {
    dayOfWeek := weekDay(2)

    fmt.Println(dayOfWeek)
}

func weekDay(number int) string {
    switch number {
        case 1:
            return "Sunday"
        case 2:
            return "Monday"
        case 3:
            return "Tuesday"
        case 4:
            return "Wednesday"
        case 5:
            return "Thursday"
        case 6:
            return "Friday"
        case 7:
            return "Saturday"
        default:
            return "Invalid Number"
    }
}
