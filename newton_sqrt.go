package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	check := 0.0
	for i:=0; i<=100; i++ {
		z = z - (z * z - x)/(2*z)
		fmt.Println("Try #", i, "Newton thinks that SQRT = ", z)
		if i != 0 {
			if check - z < 0.000000000000000000000000000000001 {
				fmt.Println("Value is pretty accurate")
					break
			}
		}
		check = z
	}
	return z
}

func main() {
	fmt.Println("SQRT of 2 is ", Sqrt(2))
}
