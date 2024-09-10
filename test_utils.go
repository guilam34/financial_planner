package main

import (
	"math"
	"strconv"
)

const float64EqualityThreshold = 1

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func fromScientificNotation(a string) float64 {
	val, _ := strconv.ParseFloat(a, 64)
	return val
}
