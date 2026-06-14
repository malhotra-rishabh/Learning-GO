package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%d Number is even\n", i)
		} else {
			fmt.Printf("%d Number is odd\n", i)
		}
	}
}
