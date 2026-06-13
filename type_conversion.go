package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 10
	var b float32 = float32(a)
	var c bool = bool(a == 10)
	var d string = string(a)             // a will act as an ASCII value and will be converted to its corresponding character which is \n
	var e string = strconv.Itoa(a)       // used to convert int to string
	var f string = strconv.FormatBool(c) // used to convert bool to string
	var g string = strconv.FormatFloat(98.2324239872632, 'f', 2, 64)

	/**
	String have 3 types of quotes:
	1. Double quotes: "Hello World" - used for string literals
	2. Backticks: `Hello World` - used for raw string literals
	3. Single quotes: 'H' - used for rune literals, meaning here h is ASCII respresentation of the character h
	**/
	fmt.Println(a, b, c, d, e, f, g)
}
