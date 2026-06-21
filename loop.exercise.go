package main

import (
	"fmt"
	"math"
)

func round(x float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return math.Round(x*pow) / pow
}

// https://en.wikipedia.org/wiki/Newton%27s_method for finding square root of a number
func sqrt(x float64, precision int) float64 {
	largerGuess := 1.0
	for largerGuess*largerGuess < x {
		largerGuess += 1
	}

	smallerGuess := largerGuess - 1
	guess := (largerGuess + smallerGuess) / 2
	guess = round(guess, precision)

	for {
		ans := 0.5 * (guess + (x / guess))
		ans = round(ans, precision)

		if ans == guess {
			break
		}

		guess = ans
	}

	return guess
}

func main() {
	fmt.Println(sqrt(200, 2))
}
