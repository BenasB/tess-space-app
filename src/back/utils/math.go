package utils

import (
	"cmp"
	"slices"
)

func GetMinMaxPercentiles(values []float32, minPercentile, maxPercentile float32) (float32, float32) {
	sortedData := make([]float32, len(values))
	copy(sortedData, values)
	slices.Sort(sortedData)
	vmin := sortedData[int(float32(len(sortedData)-1)*minPercentile/100.0)]
	vmax := sortedData[int(float32(len(sortedData)-1)*maxPercentile/100.0)]
	return vmin, vmax
}

func Clamp[T cmp.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
