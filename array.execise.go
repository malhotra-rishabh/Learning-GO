package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8 = make([][]uint8, dy)
	var count = 0
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			count++
			result[i][j] = uint8(count)
		}
	}

	return result
}

/**
Using range
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	count := uint8(0)

	for i := range result {
		result[i] = make([]uint8, dx)

		for j := range result[i] {
			count++
			result[i][j] = count
		}
	}

	return result
}

**/

func main() {
	fmt.Println(Pic(4, 3))
}
