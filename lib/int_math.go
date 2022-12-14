package lib

import (
	"math"
)

func PowInt(base, power int) int {
	if power == 0 {
		return 1
	}
	ret := base
	for i := 0; i < power; i++ {
		ret *= base
	}
	return ret
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func DivideRoundingUp(x, divisor int) int {
	return (x + divisor-1) / divisor
}

func SqDistInt(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}

func MinInt(args ...int) int {
	minValue := math.MaxInt
	for i := range args {
		if args[i] < minValue {
			minValue = args[i]
		}
	}
	return minValue
}
