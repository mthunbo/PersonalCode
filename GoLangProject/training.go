package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for z1 := 0.0; z1 != z; {
		z1 = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(81))
}
