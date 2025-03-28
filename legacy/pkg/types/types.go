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

type Library struct {
	// QuantizationCodes ??
	ResidualVectors []Point
	HashFunctions   []Point
	HashCodes       [][]int
}

// FSPointDist2Q represents a point and its distance to a query point used for full sort
type FSPointDist2Q struct {
	Point Point
	Dist  float32
}

// L0Index represents the output of KMeans clustering
type L0Index struct {
	Centroid2Points map[int][]int
	Centroids       []Point
	Point2Centroid  map[int]int
	ResidualVectors []Point
}

type LLIndex struct {
	Levels []L0Index
}

// Dist2H calculates distance from point to hyperplane
func (h *Hyperplane) Dist2H(p *Point) float32 {
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

// Dist2P calculates distance from point to another point
func (p1 *Point) Dist2P(p2 *Point) float32 {
	sum := float32(0.0)
	for j, c1 := range p1.Coordinates {
		diff := c1 - p2.Coordinates[j]
		sum += diff * diff
	}
	return float32(math32.Sqrt(sum))
}

func (p *Point) L2norm() float32 {
	sum := float32(0.0)
	for _, d := range p.Coordinates {
		sum += d * d
	}
	return math32.Sqrt(sum)
}

func (p1 *Point) Ip(p2 *Point) float32 {
	sum := float32(0.0)
	coords2 := p2.Coordinates
	for i, coord1 := range p1.Coordinates {
		sum += coord1 * coords2[i]
	}
	return sum
}

/*
   Below is the priority queue implementation for the priority queue based search
*/

// PointDist2Q represents a point and its distance to a query in the priority queue
type PQPointDist2Q struct {
	ID    int
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

// Vi skal finde ud af hvordan vi egentlig forstår vores data/-strukturer. Fx har vi ikke rigtig nogen definition af hvad et
// codeword/en codebook er. Ligenu er det som om det hele bliver lidt et miskmask af structs og slices, så måske vi skulle prøve at
// tage den top-down, og sige hvad skal der returneres i sidste ende og så bryde det ned derfra.
//
