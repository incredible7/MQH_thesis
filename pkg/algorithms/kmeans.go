package algorithms

import (
	"MQH_THESIS/pkg/types"
)

func KMeans(data []types.Point, d int, k int, n int) (map[int][]types.Point, []types.Point) {

	centroids := make([]types.Point, k)
	copy(centroids, data[0:k])

	m := make(map[int][]types.Point)

	for i := 0; i < n; i++ {
		minDist := centroids[0].Dist2P(&data[i])
		minIndex := 0
		for j := 1; j < k; j++ {
			dist := centroids[j].Dist2P(&data[i])
			if dist < minDist {
				minDist = dist
				minIndex = j
			}
		}
		m[minIndex] = append(m[minIndex], data[i])
	}

	return m, centroids
}
