package main // signifies file is main

import (
	"fmt" // Double quotes are used in Go (fmt is responsible for formatting and printing)
	"math/rand"
)

func main() {
	fmt.Println("Hello World !!", "My name is rishabh")
	fmt.Println("Random number: ", rand.Intn(10))
}
