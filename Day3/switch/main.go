package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	switch today {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Time to work!")
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("hi", today)
	}
}
