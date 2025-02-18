package helpers

import (
	"MQH_THESIS/pkg/types"
)

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

// func MergeSubtroids(l []types.L0Index, d int) types.L0Index {
// 	mergedCentroids := make([]types.Point, 256)
// 	mergedP2C := make(map[int]int)
// 	mergedC2Ps := make(map[int][]int)

// 	for i := 0; i < len(data); i++ {
		
// 		for _,index := range l {
// 			for p,c := range index.Point2Centroid {
				
// 			}
// 		}

// 		mergedCoords := make([]float32, d);
// 		k,v := 

// 	}



// }
