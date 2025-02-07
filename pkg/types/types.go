package types

import (
	"github.com/chewxy/math32"
)

// Point is defined here because it is used in the PQPointDist2Q struct
type Point struct {
	ID          int
	Coordinates []float32
}

// Hyperplane represents a hyperplane in the dataset
type Hyperplane struct {
	Q []float32
	B float32
}

// FSPointDist2Q represents a point and its distance to a query point used for full sort
type FSPointDist2Q struct {
	Point Point
	Dist  float32
}

// Dist calculates distance from point to hyperplane
func (h *Hyperplane) Dist(p *Point) float32 {
	numerator := h.B
	for i := range p.Coordinates {
		numerator = numerator + h.Q[i]*p.Coordinates[i]
	}
	if numerator < 0 {
		numerator = -numerator
	}
	denominator := float32(0.0)
	for i := range h.Q {
		denominator += h.Q[i] * h.Q[i]
	}
	denominator = float32(math32.Sqrt(denominator))
	return numerator / denominator
}
