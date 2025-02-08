package algorithms

import (
	"MQH_THESIS/pkg/types"
	"fmt"
)

func alg1(data []types.Point, d int, coarse_centroids_number int, L int) {
	coarse_quant, coarse_centroids := kmeans(data, d, 256, 100)

	res_vecs := make([][]types.Point, L+1)
	res_vecs[0] = 

}

func residual_calc(data []types.Point, quantizations map[int][]int, centroids []types.Point) []types.Point {
	residuals := make([]types.Point, len(data))
	for CID, ps := range quantizations {
		centroid_coords := centroids[CID].Coordinates
		for _, PID := range ps {
			point_coords := data[PID].Coordinates
			res_coords := make([]float32, len(point_coords))
			for i := 0; i < len(point_coords); i++ {
				res_coords[i] = point_coords[i] - centroid_coords[i]
			}
			residuals[PID] = types.Point{
				ID:          PID,
				Coordinates: res_coords,
			}
		}
	}
	return residuals
}
