package base

import "math"

const (
	BEIJING = iota
	SHANGHAI
	JIANGSU
	CHONGQING
)

func City() uint {
	return BEIJING
}

func Sqrt(x float64) float64 {
	switch {
	case x == 0:
		return x
	case x < 0:
		return math.NaN()
	}
	z := float64(1)
	z -= (z*z - x) / (2*z)
	return z
}
