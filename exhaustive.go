/*
Set up currently to run on the Cifar-10 dataset only.
*/

package main

import "fmt"
import "math"

func main() {
	// dataset := ...
	// Query:= ...

}

type Point struct {
	ID          int
	Coordinates []float64
}

type Hyperplane struct {
	Q []float64
	B float64
}

func (h *Hyperplane) dist(p *Point) float64 {
	numerator := h.B
	for i := range p.Coordinates {
		numerator = numerator + h.Q[i]*p.Coordinates[i]
	}
	numerator = math.Abs(numerator)
	denominator := 0.0
	for i := range h.Q {
		denominator = denominator + h.Q[i]*h.Q[i]
	}
	denominator = math.Sqrt(denominator)
	return numerator / denominator
}
