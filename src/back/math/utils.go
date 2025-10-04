package math

import "slices"

func GetMinMaxPercentiles(values []float32, minPercentile, maxPercentile float32) (float32, float32) {
	sortedData := make([]float32, len(values))
	copy(sortedData, values)
	slices.Sort(sortedData)
	vmin := sortedData[int(float32(len(sortedData)-1)*minPercentile/100.0)]
	vmax := sortedData[int(float32(len(sortedData)-1)*maxPercentile/100.0)]
	return vmin, vmax
}
