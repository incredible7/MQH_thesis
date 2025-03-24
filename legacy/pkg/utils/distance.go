package utils

func P2H_dist(point []float32, normal []float32, b float32) float32 {
	numerator := b
	denominator := float32(0.0)
	for i := range point {
		numerator += point[i] * normal[i]
		denominator += normal[i] * normal[i]
	}
	return numerator / denominator
}