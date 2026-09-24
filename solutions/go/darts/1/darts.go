package darts

import (
	"math"
)

func isOnTarget(x, y, centerX, centerY, radius float64) bool {
	return math.Pow((x-centerX), 2)+math.Pow((y-centerY), 2) <= math.Pow(radius, 2)
}

func Score(x, y float64) int {
	if isOnTarget(x, y, 0, 0, 1) {
		return 10
	} else if isOnTarget(x, y, 0, 0, 5) {
		return 5
	} else if isOnTarget(x, y, 0, 0, 10) {
		return 1
	} else {
		return 0
	}
}
