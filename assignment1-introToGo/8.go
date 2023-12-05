package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	if x < 0 {
		return math.NaN()
	}
	z = 1.
	i := 1
	for i < 10 {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		i++
	}
	fmt.Println("Finished Sqrt")

	return
}

func Sqrt2(x float64) (z float64) {
	if x < 0 {
		return math.NaN()
	}
	z = x
	prev := 0.
	for math.Abs(z-prev) > 1e-9 {
		fmt.Println(z)
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	fmt.Println("Finished Sqrt2")

	return
}

func main() {

	fmt.Println(Sqrt(2.), Sqrt2(2.), math.Sqrt(2.))

}
