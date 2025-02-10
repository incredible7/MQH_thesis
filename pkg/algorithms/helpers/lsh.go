package helpers

import (
	"MQH_THESIS/pkg/types"
	"math/rand"
)

func LSH(residuals []types.Point, m int, d int) ([]types.Point, [][]bool) {
	n := len(residuals)
	alphas := generateAlphas(m, d)
	hashvalues := make([][]bool, n)
	for i := 0; i < n; i++ {
		hashvalues[i] = generateBitstring(residuals[i], alphas)
	}
	return alphas, hashvalues
}

func generateAlphas(m int, d int) []types.Point {
	alphas := make([]types.Point, m)
	for i := 0; i < m; i++ {
		coords := make([]float32, d)
		for j := 0; j < d; j++ {
			coords[j] = randomFloat32()
		}
		alphas[i] = types.Point{
			ID:          i,
			Coordinates: coords,
		}
	}
	return alphas
}

func generateBitstring(p types.Point, alphas []types.Point) []bool {
	bitstring := make([]bool, len(alphas))
	for i, alpha := range alphas {
		if p.Ip(alpha) >= 0 {
			bitstring[i] = true
		} else {
			bitstring[i] = false
		}
	}
	return bitstring
}

func randomFloat32() float32 {
	return -1.0 + rand.Float32()*(2.0)
}
