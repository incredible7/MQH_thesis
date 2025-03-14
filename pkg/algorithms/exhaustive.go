package algorithms

import (
	"MQH_THESIS/pkg/types"
	"MQH_THESIS/pkg/algorithms/utils"
	"container/heap"
	"fmt"
	"slices"
	"os"
	"sort"
	"time"
)
// creates a index for the points and returns a pointer to it
func preprocess(points []float32, n int, d int) {
		points = points
	}

func search(points []float32, normal []float32, b float32, n int, d int, k int, MinKList []int) {
	// create a priority queue
	pq := make(types.DistancePriorityQueue, 0)
	heap.Init(&pq)

	// add all points to the priority queue
	for i := 0; i < n; i++ {
		point := points[i*d : (i+1)*d]
		dist := utils.P2H_dist(point, normal, b)
		itempointer := &types.PQPointDist2Q{
			ID: i,
			Dist:  dist,
		}
		heap.Push(&pq, itempointer)
	}
	// get the k nearest neighbors
	for i := 0; i < k; i++ {
		// directly assign the ID to the MinKList since we are using a pointer
		MinKList[i] = heap.Pop(&pq).(*types.PQPointDist2Q).ID
	}
}



// // ExhaustiveFS performs full sort search
// func ExhaustiveFS(dataset string, points []types.Point, hyperplanes []types.Hyperplane, nq int, k int, suffix string) {
// 	// create a file to write the results to
// 	outfile, err := os.Create("data/results/" + dataset + ".fs" + suffix)
// 	if err != nil {
// 		fmt.Println("Error creating results file:", err)
// 		return
// 	}
// 	defer outfile.Close()

// 	// Write header - now with flexible nq and k
// 	fmt.Fprintf(outfile, "%d,%d\n", nq, k)

// 	// use timer from library
// 	start := time.Now()

// 	// For each hyperplane check all points and their distance to the hyperplane
// 	for i := 0; i < nq; i++ {
// 		fmt.Fprintf(outfile, "%d", i+1)

// 		// Sort all points by distance to this hyperplane
// 		allPoints := make([]types.FSPointDist2Q, len(points))

// 		for j := range points {
// 			allPoints[j].Point = points[j]
// 			allPoints[j].Dist = hyperplanes[i].Dist2H(&points[j])
// 		}

// 		sort.Slice(allPoints, func(a, b int) bool {
// 			return allPoints[a].Dist < allPoints[b].Dist
// 		})

// 		// Write k nearest neighbors (instead of fixed 100)
// 		for j := 0; j < k; j++ {
// 			fmt.Fprintf(outfile, ",%d,%.9f", allPoints[j].Point.ID+1, allPoints[j].Dist)
// 		}
// 		fmt.Fprintln(outfile)
// 	}

// 	elapsed := time.Since(start)
// 	fmt.Printf("FS took %s\n", elapsed)
// }

// // ExhaustivePQ performs priority queue based search
// func ExhaustivePQ(dataset string, points []types.Point, hyperplanes []types.Hyperplane, nq, k int, suffix string) {
// 	// create a file to write the results to
// 	outfile, err := os.Create("data/results/" + dataset + ".pq" + suffix)
// 	if err != nil {
// 		fmt.Println("Error creating results file:", err)
// 		return
// 	}
// 	defer outfile.Close()

// 	// Write header with flexible nq and k
// 	fmt.Fprintf(outfile, "%d,%d\n", nq, k)

// 	//use timer from library
// 	start := time.Now()

// 	// For each hyperplane create a priority queue and add all points to it with their distance to that hyperplane
// 	for i := 0; i < nq; i++ {
// 		fmt.Fprintf(outfile, "%d", i+1)

// 		// Create priority queue
// 		pq := make(types.DistancePriorityQueue, 0)
// 		heap.Init(&pq)

// 		// Add all points to priority queue
// 		for j := range points {
// 			dist := hyperplanes[i].Dist2H(&points[j])
// 			item := &types.PQPointDist2Q{
// 				Point: points[j],
// 				Dist:  dist,
// 			}
// 			heap.Push(&pq, item)
// 		}

// 		// Get k nearest neighbors (instead of fixed 100)
// 		for j := 0; j < k; j++ {
// 			item := heap.Pop(&pq).(*types.PQPointDist2Q)
// 			fmt.Fprintf(outfile, ",%d,%.9f", item.Point.ID+1, item.Dist)
// 		}
// 		fmt.Fprintln(outfile)
// 	}
// 	elapsed := time.Since(start)
// 	fmt.Printf("PQ took %s\n", elapsed)
// }
