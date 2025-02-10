package helpers

import (
	"MQH_THESIS/pkg/types"
	"math/rand"
)

func LSH(residuals []types.Point, m int, d int) [][]int {
	n := len(residuals)
	alphas := generateAlphas(m, d)
	hashvalues := make([][]int, n)
	for i := 0; i < n; i++ {
		hashvalues[i] = generateBitstring(residuals[i], alphas)
	}
	return hashvalues
}

func generateAlphas(m int, d int) []types.Point {
	alphas := make([]types.Point, m)
	for i := 0; i < m; i++ {
		coords := make([]float32, d)
		for j := 0; j < d; j++ {
			coords[j] = rand.Float32()
		}
		alphas[i] = types.Point{
			ID:          i,
			Coordinates: coords,
		}
	}
	return alphas
}

func generateBitstring(p types.Point, alphas []types.Point) []int {
	bitstring := make([]int, len(alphas))
	for i, alpha := range alphas {
		if p.Ip(alpha) >= 0 {
			bitstring[i] = 1
		} else {
			bitstring[i] = 0
		}
	}
	return bitstring
}
