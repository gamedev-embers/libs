package mathutils

import "math"

// f(x) = ax^3 + bx^2 + cx + d
func FormulaX3(arr []uint32, x uint32) uint32 {
	_ = arr[3]
	if x == 0 {
		return arr[3]
	} else if x == 1 {
		return arr[0] + arr[1] + arr[2] + arr[3]
	}
	square := x * x
	cube := square * x
	return arr[0]*cube + arr[1]*square + arr[2]*x + arr[3]
}

// only for performance reference
func formulaX3_1Pow(arr []uint32, x uint32) uint32 {
	_v := uint32(math.Pow(float64(x), 2))
	return arr[0]*_v*x + arr[1]*_v + arr[2]*x + arr[3]
}

// only for performance reference
func formulaX3_2Pow(arr []uint32, x uint32) uint32 {
	_x := float64(x)
	ax3 := arr[0] * uint32(math.Pow(_x, 3))
	bx2 := arr[1] * uint32(math.Pow(_x, 2))
	cx1 := arr[2] * x
	d := arr[3]
	return ax3 + bx2 + cx1 + d
}
