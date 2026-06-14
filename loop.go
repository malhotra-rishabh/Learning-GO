package main

import "fmt"

// go only have for loop
func main() {
	// normal for loop syntax
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// but can be used as while
	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Sum: ", sum)
}
