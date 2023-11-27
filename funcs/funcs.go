package funcs

import "math"

func MinValues(num1 float64, num2 float64, num3 float64) float64 {
	compareTwo := math.Min(num1, num2)
	result := math.Min(compareTwo, num3)
	return result
}
func AvgValues(num1 float64, num2 float64, num3 float64) float64 {
	result := (num1 + num2 + num3) / 3
	return result
}
func DiffEq(num1 float64, num2 float64) float64 {
	result := (num1 + num2) / num1
	return result
}
