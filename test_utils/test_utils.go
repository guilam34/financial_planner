package test_utils

import (
	"math"
)

const float64EqualityThreshold = 1

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
