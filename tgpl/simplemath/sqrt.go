package simplemath

import "math"

// Sqrt for i
func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}
