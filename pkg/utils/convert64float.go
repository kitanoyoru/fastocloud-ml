package utils

func ConvertToFloat64Slice(arr []float32) []float64 {
	newarr := make([]float64, len(arr))
	for i, v := range arr {
		newarr[i] = float64(v)
	}

	return newarr
}
