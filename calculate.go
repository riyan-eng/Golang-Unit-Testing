package golangunittesting

import "math"

func CalculateIsArmstrong(number int) bool {
	a := number / 100
	b := number % 100 / 10
	c := number % 10
	return number == int(math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3))
}
