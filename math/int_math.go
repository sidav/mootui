package math

import "math"

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
