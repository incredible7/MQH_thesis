package algorithms

import (
	"MQH_THESIS/pkg/types"
	"slices"
)

func kmeans(data []types.Point, d int, k int, iterations int) (map[int][]int, []types.Point) {
	// initialize centroids
	centroids := make([]types.Point, k)
	copy(centroids, data[:k])
	m := make(map[int][]int)

	for i := 0; i < iterations; i++ {
		m = assign(data, centroids)
		centroids = update(data, m, d, k)
	}
	m = assign(data, centroids)
	return m, centroids
}

func assign(data []types.Point, centroids []types.Point) map[int][]int {
	m := make(map[int][]int)

	for i := range data {
		// min := float32(math32.MaxFloat32)
		dists := make([]float32, len(centroids))
		for j := range centroids {
			dist := (data)[i].Dist2P(&centroids[j])
			dists[j] = dist
		}
		// if dist < min {
		// 	min = dist
		// 	if cid == -1 {
		// 		old := m[cid]
		// 		(old, data[i].ID)
		// 	}
		// 	cid = centroids[j].ID

		min := slices.Min(dists)
		index := slices.Index(dists, min)
		m[index] = append(m[index], (data)[i].ID)
	}
	return m
}

func update(data []types.Point, m map[int][]int, d int, k int) []types.Point {
	new_centroids := make([]types.Point, k)
	for i, points := range m {
		new_coords := make([]float32, d)
		for _, PID := range points {
			for h := range d {
				new_coords[h] += data[PID].Coordinates[h]
			}
		}
		new_coords = Map(new_coords, func(x float32) float32 { return x / float32(len(points)) })
		new_centroids[i] = types.Point{
			ID:          i,
			Coordinates: new_coords}
	}
	return new_centroids
}

func Map(s []float32, fn func(float32) float32) []float32 {
	result := make([]float32, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}
