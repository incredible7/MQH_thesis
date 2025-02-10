package algorithms

import (
	"MQH_THESIS/pkg/types"
	"slices"
)

// L0Index represents the output of KMeans clustering
// type L0Index struct {
// 	Centroid2Points map[int][]int
// 	Centroids       []Point
// 	Point2Centroid  map[int]int
// 	ResidualVectors []Point
// }

func ProductPartitioning(data []types.Point, d int, k int, iterations int, m int) []types.L0Index {
	fullData := make([]types.L0Index, m)
	divResult := d / m

	// For each sub-space (partition of dimensions)
	for i := 0; i < m; i++ {
		// Here we create a new array of points, each point only has the dimensions for this subspace.
		subspacePoints := make([]types.Point, len(data))
		start := i * divResult
		end := start + divResult
		// if i = m-1 then this must be the last subspace
		if i == m-1 {
			end = d
		}

		// Copy points over with only the dimensions for this subspace
		for j, point := range data {
			subspacePoints[j] = types.Point{
				ID:          point.ID,
				Coordinates: point.Coordinates[start:end],
			}
		}

		// Run k-means on this subspace
		fullData[i] = KMeans(subspacePoints, end-start, k, iterations)
	}

	// Run some function here to merge the L0Index structs into a single L0Index struct?

	return fullData
}

// KMeans performs the k-means clustering algorithm on the given data points and returns L0Result struct.
func KMeans(data []types.Point, d int, k int, iterations int) types.L0Index {
	centroids := make([]types.Point, k)
	copy(centroids, data[:k])

	centroid2Points := make(map[int][]int)

	for i := 0; i < iterations; i++ {
		centroid2Points = assign(data, centroids)
		centroids = update(data, centroid2Points, d, k)
	}

	point2Centroid := make(map[int]int)

	// also map from p2c as we do for c2p's..
	for centroidID, points := range centroid2Points {
		for _, pointID := range points {
			point2Centroid[pointID] = centroidID
		}
	}

	return types.L0Index{
		Centroid2Points: centroid2Points,
		Centroids:       centroids,
		Point2Centroid:  point2Centroid,
		ResidualVectors: make([]types.Point, len(data)),
	}
}

// assign assigns each data point to the nearest centroid.
func assign(data []types.Point, centroids []types.Point) map[int][]int {
	m := make(map[int][]int)

	for i := range data {
		dists := make([]float32, len(centroids))
		for j := range centroids {
			dist := (data)[i].Dist2P(&centroids[j])
			dists[j] = dist
		}
		min := slices.Min(dists)
		index := slices.Index(dists, min)
		m[index] = append(m[index], (data)[i].ID)
	}
	return m
}

// update updates the centroids based on the assigned data points.
func update(data []types.Point, m map[int][]int, d int, k int) []types.Point {
	newCentroids := make([]types.Point, k)
	for i, points := range m {
		newCoords := make([]float32, d)
		for _, PID := range points {
			for h := range d {
				newCoords[h] += data[PID].Coordinates[h]
			}
		}
		// Calculate the mean of the coordinates using the Map function which applies the division to each element of the slice.
		newCoords = funcMap(newCoords, func(x float32) float32 { return x / float32(len(points)) })
		newCentroids[i] = types.Point{
			ID:          i,
			Coordinates: newCoords}
	}
	return newCentroids
}

// funcMap applies the given function to each element of the slice.
func funcMap(s []float32, fn func(float32) float32) []float32 {
	result := make([]float32, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}
