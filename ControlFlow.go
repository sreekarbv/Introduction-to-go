package main

import "fmt"

func main() {
	// If-else statement
	x := 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than or equal to 10")
	}

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	day := "Tuesday"
	// Switch statement
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday", "Wednesday":
		fmt.Println("It's Tuesday or Wednesday")
	default:
		fmt.Println("It's another day")
	}
}
