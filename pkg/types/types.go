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

/*
   Below is the priority queue implementation for the priority queue based search
*/

// PointDist2Q represents a point and its distance to a query in the priority queue
type PQPointDist2Q struct {
	Point Point
	Dist  float32
	Index int // Required by heap.Interface
}

// A DistancePriorityQueue implements heap.Interface and holds PointDist2Qs
type DistancePriorityQueue []*PQPointDist2Q

func (pq DistancePriorityQueue) Len() int { return len(pq) }

// The Less functuon is our min queue, so we want the smallest distances first
func (pq DistancePriorityQueue) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

func (pq DistancePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *DistancePriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PQPointDist2Q)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *DistancePriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}
