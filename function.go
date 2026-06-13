package main

import "fmt"

// normal way to define a function
func add(a int, b int) int {
	return a + b
}

// shorthand way to define a function arguments of same type
func addShortHand(a, b int) int {
	return a + b
}

// multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println("Sum: ", add(4, 5))
	fmt.Println("Sum shorthand: ", addShortHand(2, 1))

	a := "hello"
	b := "world" // := is only used for variable declaration and assignment in one line, otherwise we have to use = for assignment after declaration
	a, b = swap(a, b)
	fmt.Println("After swap: ", a, b)
}
