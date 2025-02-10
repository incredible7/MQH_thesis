package algorithms

import (
	"MQH_THESIS/pkg/algorithms/helpers"
	"MQH_THESIS/pkg/types"
	"fmt"
	"time"
)

// Mqh ...
func Mqh(dataset string, points []types.Point, hyperplanes []types.Hyperplane, nq int, d int, n int, suffix string, levels int) {
	fmt.Println("Starting MQH")
	L0result := coarseQuantization(points, d, n)

	multiLevelQuantization(L0result.ResidualVectors, d, n)
}

func coarseQuantization(points []types.Point, d int, n int) types.L0Index {
	// use timer from library
	start := time.Now()
	fmt.Println("Running coarse quantization")
	k := 256
	iterations := 1
	L0Result := helpers.KMeans(points, d, k, iterations)

	fmt.Printf("KMeans generated %d centroids (k=%d)\n", len(L0Result.Centroids), k)

	// Calculate residuals
	L0Result.ResidualVectors = calculateResiduals(points, L0Result.Centroids, L0Result.Point2Centroid)

	elapsed := time.Since(start)
	fmt.Printf("Coarse quantization on %d points for %d iterations took %s\n", n, iterations, elapsed)

	return L0Result
}

// calculateResiduals calculates the residual vectors for all points
func calculateResiduals(data []types.Point, centroids []types.Point, point2Centroid map[int]int) []types.Point {
	residualVectors := make([]types.Point, len(data))

	for i := range data {
		residualVectors[i] = types.Point{
			ID:          data[i].ID,
			Coordinates: make([]float32, len(data[i].Coordinates)),
		}

		centroidIndex := point2Centroid[data[i].ID]
		for j := range data[i].Coordinates {
			residualVectors[i].Coordinates[j] = data[i].Coordinates[j] - centroids[centroidIndex].Coordinates[j]
		}
	}
	return residualVectors
}

func multiLevelQuantization(points []types.Point, d int, n int) {

}
