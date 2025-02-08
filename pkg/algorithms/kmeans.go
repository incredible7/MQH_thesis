package algorithms

import (
	"MQH_THESIS/pkg/types"
	"slices"
)

// ProductKMeans is in charge of splitting the data into c codebooks and then running k-means on each codebook.
// It returns a slice of CodebookData, where each CodebookData represents a codebook and contains the centroid ID as the key and the point IDs as the value.
func ProductKMeans(data []types.Point, d int, k int, iterations int, c int) []types.Codebook {
	dimPerBook := d / c
	codebooks := make([]types.Codebook, c)

	for cb := range codebooks {
		// Calculate dimension range for this codebook
		start := cb * dimPerBook
		end := start + dimPerBook
		if cb == c-1 {
			end = d // for last codebook we need to just take the rest of the d
		}

		// Create subset of points with only the relevant dimensions
		subset := make([]types.Point, len(data))
		for i, p := range data {
			subset[i] = types.Point{
				ID:          p.ID,
				Coordinates: p.Coordinates[start:end],
			}
		}

		assignments, centroids := KMeans(subset, end-start, k, iterations)
		codebooks[cb] = types.Codebook{
			Assignments: assignments,
			Centroids:   centroids,
		}
	}

	return codebooks
}

// KMeans performs the k-means clustering algorithm on the given data points.
// Parameters:
// - data: a slice of Point representing the data points.
// - d: the dimensionality of the data points.
// - k: the number of clusters.
// - iterations: the number of iterations to run the algorithm.
// Returns a map where the key is the centroid index and the value is a slice of data point IDs assigned to that centroid.
func KMeans(data []types.Point, d int, k int, iterations int) (map[int][]int, []types.Point) {
	centroids := make([]types.Point, k)
	copy(centroids, data[:k])
	m := make(map[int][]int)

	for i := 0; i < iterations; i++ {
		m = assign(data, centroids)
		centroids = update(data, m, d, k)
	}
	return m, centroids
}

// assign assigns each data point to the nearest centroid.
// Parameters:
// - data: a slice of Point representing the data points.
// - centroids: a slice of Point representing the centroids.
// Returns a map where the key is the centroid index and the value is a slice of data point IDs assigned to that centroid.
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
// Parameters:
// - data: a slice of Point representing the data points.
// - m: a map where the key is the centroid index and the value is a slice of data point IDs assigned to that centroid.
// - d: the dimensionality of the data points.
// - k: the number of clusters.
// Returns a slice of Point representing the new centroids.
func update(data []types.Point, m map[int][]int, d int, k int) []types.Point {
	new_centroids := make([]types.Point, k)
	for i, points := range m {
		new_coords := make([]float32, d)
		for _, PID := range points {
			for h := range d {
				new_coords[h] += data[PID].Coordinates[h]
			}
		}
		// Calculate the mean of the coordinates using the Map function which applies the division to each element of the slice.
		new_coords = Map(new_coords, func(x float32) float32 { return x / float32(len(points)) })
		new_centroids[i] = types.Point{
			ID:          i,
			Coordinates: new_coords}
	}
	return new_centroids
}

// Map applies the given function to each element of the slice.
// Parameters:
// - s: a slice of float32.
// - fn: a function that takes a float32 and returns a float32.
// Returns a new slice where each element is the result of applying the function to the corresponding element in the input slice.
func Map(s []float32, fn func(float32) float32) []float32 {
	result := make([]float32, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}
