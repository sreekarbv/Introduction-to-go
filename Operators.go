package main

import "fmt"

func main() {
	a := 10
	b := 5

	// Arithmetic operators
	fmt.Println(a + b) // Addition
	fmt.Println(a - b) // Subtraction
	fmt.Println(a * b) // Multiplication
	fmt.Println(a / b) // Division
	fmt.Println(a % b) // Modulus

	// Comparison operators
	fmt.Println(a == b) // Equal to
	fmt.Println(a != b) // Not equal to
	fmt.Println(a > b)  // Greater than
	fmt.Println(a < b)  // Less than
	fmt.Println(a >= b) // Greater than or equal to
	fmt.Println(a <= b) // Less than or equal to

	// Logical operators
	fmt.Println(a > 5 && b < 10) // AND
	fmt.Println(a > 5 || b < 10) // OR
	fmt.Println(!(a > 5))        // NOT
}
