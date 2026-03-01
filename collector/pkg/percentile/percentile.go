package percentile

import "sort"

func P(values []float64, percentile float64) float64 {
	if len(values) == 0 {
		return 0
	}
	if percentile <= 0 {
		return min(values)
	}
	if percentile >= 100 {
		return max(values)
	}

	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	idx := int((percentile / 100.0) * float64(len(sorted)-1))
	return sorted[idx]
}

func min(values []float64) float64 {
	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func max(values []float64) float64 {
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}
