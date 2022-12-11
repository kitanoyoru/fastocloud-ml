package utils

import "math"

func GetMinNumInSlice(arr []float64) (float64, int) {
	num, idx := math.MaxFloat64, 0
	for i, v := range arr {
		if v < num {
			num = v
			idx = i
		}
	}

	return num, idx
}
