package helpers

import (
	"MQH_THESIS/pkg/types"
)

func NERQ(data []types.Point, m int) []types.Point {
	normalized := normalize(data)
	return normalized
}

func normalize(data []types.Point) []types.Point {
	normalized := make([]types.Point, len(data))
	for i, p := range data {
		norm := p.L2norm()
		normalizedCoords := make([]float32, len(p.Coordinates))
		for j, c := range p.Coordinates {
			normalizedCoords[j] = c / norm
		}
		normalized[i] = types.Point{
			ID:          i,
			Coordinates: normalizedCoords,
		}
	}
	return normalized
}
