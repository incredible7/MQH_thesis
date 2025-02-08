package algorithms

import (
	"MQH_THESIS/pkg/types"
	"fmt"
	"os"
	"time"
)

// Mqh ...
func Mqh(dataset string, points []types.Point, hyperplanes []types.Hyperplane, nq int, d int, k int, n int, suffix string, levels int) {
	// create a file to write the results to
	outfile, err := os.Create("data/results/" + dataset + ".fs" + suffix)
	if err != nil {
		fmt.Println("Error creating results file:", err)
		return
	}
	defer outfile.Close()

	// Write header - now with flexible nq and k
	fmt.Fprintf(outfile, "%d,%d\n", nq, k)

	// use timer from library
	start := time.Now()
	amntcodebooks := int(2)

	codebooks := ProductKMeans(points, d, k, 2, amntcodebooks)

	for i := 0; i < amntcodebooks; i++ {
		fmt.Printf("Codebook %d:\n", i)
		fmt.Printf("  Centroid coordinates: %v\n", codebooks[i].Centroids[0])
		fmt.Printf("  Assigned points: %v\n\n", codebooks[i].Assignments[0])
	}

	// We want to

	// // For each hyperplane check all points and their distance to the hyperplane
	// for i := 0; i < nq; i++ {
	// 	fmt.Fprintf(outfile, "%d", i+1)

	// 	// Sort all points by distance to this hyperplane
	// 	allPoints := make([]types.FSPointDist2Q, len(points))

	// 	for j := range points {
	// 		allPoints[j].Point = points[j]
	// 		allPoints[j].Dist = hyperplanes[i].Dist(&points[j])
	// 	}

	// 	sort.Slice(allPoints, func(a, b int) bool {
	// 		return allPoints[a].Dist < allPoints[b].Dist
	// 	})

	// 	// Write k nearest neighbors (instead of fixed 100)
	// 	for j := 0; j < k; j++ {
	// 		fmt.Fprintf(outfile, ",%d,%.9f", allPoints[j].Point.ID+1, allPoints[j].Dist)
	// 	}
	// 	fmt.Fprintln(outfile)
	// }

	/*
		TODO: Coarse quantization of points
		Point will be represented as:
			type Point struct {
				ID          int
				Coordinates []float32
			}
		-> Try to implement eg. github.com/parallelo-ai/kmeans

		Resulting in points clustered to 256 clusters.
		Implement something like:
			type Centroid struct {
				ID          int
				Coordinates []float32
			}

			var m map[centroid][]priorityque.Point

			And then use eg.
		allQuantizedPoints := make([]types.Point2Centroid, len(256))


		Now we need to implement NERQ. THis will be represented as a recursive structure where points are quantized through 3 levels, for each level the residuals of the previous levels are used. And initially the residuals used, for it's first level are calculated from the coarse quantization, so from the distance of each point to it's coarse centroid. For the NERQ we will want to represent every point as 16 parts, where each part is this points normalvectors d/16.
		Now one of these will be based on relative norm and the other 15 on angle in each level.
		Now in each level we need to do codebook training, then quantization and then finally we want to use LSH so as to be able to represent each point through an m-bit code to easily be able to compare points later in the search algorithm based on hamming distance. Possible LSH library: https://github.com/ekzhu/lsh
		We also need to store quantization codes and hashing codes in each level in a proper structure.


	*/

	elapsed := time.Since(start)
	fmt.Printf("FS took %s\n", elapsed)
}
