package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello, " + name + "!")
}

func add(x, y int) int {
	return x + y
}

func main() {
	greet("Alice")
	result := add(5, 3)
	fmt.Println(result)
}
