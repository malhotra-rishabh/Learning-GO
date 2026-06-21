/*
*
A struct is a collection of fields.
*
*/
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	vertexObj := Vertex{1, 2}
	fmt.Println(vertexObj.X)
	fmt.Println(vertexObj.Y)
}
