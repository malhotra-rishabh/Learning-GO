package main

import "fmt"

var hello, world string

/*
*
Since the variable name "hello" is same in local and global scope, the local variable will be used in the function.
The global variable will be shadowed by the local variable.
*
*/
func main() {
	var hello, funcWorld string

	hello, world = "hello", "world"
	fmt.Println(hello, world)

	hello, funcWorld = "func hello", "func world"
	fmt.Println(hello, funcWorld)
}
