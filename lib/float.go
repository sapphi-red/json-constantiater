package lib

import (
	"math"
	"strconv"
)

//go:nosplit
func AppendFloat32(dest []byte, src float32, prec int) []byte {
	src64 := float64(src)
	abs := math.Abs(src64)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if abs != 0 {
		f32abs := float32(abs)
		if f32abs < 1e-6 || f32abs >= 1e21 {
			fmt = 'e'
		}
	}
	return strconv.AppendFloat(dest, src64, fmt, prec, 32)
}

//go:nosplit
func AppendFloat64(dest []byte, src float64, prec int) []byte {
	abs := math.Abs(src)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if abs != 0 {
		if abs < 1e-6 || abs >= 1e21 {
			fmt = 'e'
		}
	}
	return strconv.AppendFloat(dest, src, fmt, prec, 64)
}
