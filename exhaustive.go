package main

import (
	"MQH_THESIS/pkg/priorityqueue" // Wanted to show how to structure packages. PQ inspired from go docs.
	"container/heap"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/chewxy/math32"
)

// Hyperplane is a struct that represents a hyperplane in the dataset
type Hyperplane struct {
	Q []float32
	B float32
}

// FSPointDist2Q is a struct that represents a point and its distance to a query point used for full sort
type FSPointDist2Q struct {
	Point priorityqueue.Point
	Dist  float32
}

// main function
func main() {
	if len(os.Args) != 5 {
		fmt.Println("Error - try again using this format:\n go run . <dataset name> <num_points> <dimensionality> <num queries>")
		return
	}

	// Read command-line arguments
	dataset := os.Args[1]
	n, _ := strconv.Atoi(os.Args[2])
	d, _ := strconv.Atoi(os.Args[3])
	// k, _ := strconv.Atoi(os.Args[4]) - moved for now as we are doing GT
	nq, _ := strconv.Atoi(os.Args[4])

	pointsData, err := readBinaryFile("data/datasets/" + dataset + ".ds")
	if err != nil {
		fmt.Printf("Error reading dataset file: %v\n", err)
		os.Exit(1)
	}

	queriesData, err := readBinaryFile("data/datasets/" + dataset + ".q")
	if err != nil {
		fmt.Printf("Error reading query file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("🔍 Loading dataset: %s with %d points of %d dimensions\n", dataset, n, d)
	points := readPoints(pointsData, n, d)

	fmt.Printf("🔍 Loading %d hyperplane queries\n", nq)
	hyperplanes := readHyperplanes(queriesData, d)

	exhaustiveFS(dataset, points, hyperplanes, nq) //Remember k here at some point
	fmt.Println("FS is done")

	exhaustivePQ(dataset, points, hyperplanes, nq) //Remember k here at some point
	fmt.Println("PQ is done")
}

// this helper function is used for opening, getting file info from - and reading the binary file
func readBinaryFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("getting file info: %w", err)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	return data, nil
}

func readPoints(data []byte, n int, d int) []priorityqueue.Point {
	totalValues := d + 1 // Each entry has d float32 + 1 delimiter - the delimiter is used
	points := make([]priorityqueue.Point, n)

	for i := 0; i < n; i++ {
		startIndex := i * totalValues

		// Read feature vector (first d values)
		points[i] = priorityqueue.Point{
			ID:          i, // Assign sequential ID
			Coordinates: make([]float32, d),
		}
		for j := 0; j < d; j++ {
			bits := binary.LittleEndian.Uint32(data[(startIndex+j)*4 : (startIndex+j+1)*4])
			points[i].Coordinates[j] = math.Float32frombits(bits)
		}
		// Ignore the last value (which is always 1.0)

	}

	fmt.Printf("Loaded points 🧙‍♂️\n")
	return points
}

func readHyperplanes(querydata []byte, d int) []Hyperplane {
	hyperplanes := make([]Hyperplane, 100)
	for i := 0; i < 100; i++ {
		query := Hyperplane{
			Q: make([]float32, d),
			B: 0.0,
		}
		movePosition := i * (d + 1) * 4 // Move position in file by (d+1)*4 bytes for each hyperplane
		for j := 0; j < d; j++ {
			bits := binary.LittleEndian.Uint32(querydata[movePosition+(j*4) : movePosition+(j+1)*4])
			query.Q[j] = math.Float32frombits(bits)
		}
		Bbits := binary.LittleEndian.Uint32(querydata[movePosition+(d*4) : movePosition+(d+1)*4])
		query.B = math.Float32frombits(Bbits)
		hyperplanes[i] = query
	}
	fmt.Printf("Loaded hyperplanes 🛩️\n")
	return hyperplanes
}

// this function is used for writing the results of the full sort to a file
func exhaustiveFS(dataset string, points []priorityqueue.Point, hyperplanes []Hyperplane, nq int) {
	// create a file to write the results to
	outfile, err := os.Create("data/results/" + dataset + ".fs.gt")
	if err != nil {
		fmt.Println("Error creating results file:", err)
		return
	}
	defer outfile.Close()

	// Write header - 100 hyperplanes, 100 points
	fmt.Fprintf(outfile, "%d,%d\n", nq, nq)

	//use timer from library
	start := time.Now()

	// For each hyperplane check all points and their distance to the hyperplane
	for i := 0; i < 100; i++ {
		fmt.Fprintf(outfile, "%d", i+1)

		// Sort all points by distance to this hyperplane
		allPoints := make([]FSPointDist2Q, len(points))

		for j := range points {
			allPoints[j].Point = points[j]
			allPoints[j].Dist = hyperplanes[i].dist(&points[j])
		}

		sort.Slice(allPoints, func(a, b int) bool {
			return allPoints[a].Dist < allPoints[b].Dist
		})

		// Write 100 nearest neighbors
		for j := 0; j < 100; j++ {
			fmt.Fprintf(outfile, ",%d,%.9f", allPoints[j].Point.ID+1, allPoints[j].Dist)
		}
		fmt.Fprintln(outfile)
	}

	elapsed := time.Since(start)
	fmt.Printf("FS took %s\n", elapsed)
}

func exhaustivePQ(dataset string, points []priorityqueue.Point, hyperplanes []Hyperplane, nq int) {
	// create a file to write the results to
	outfile, err := os.Create("data/results/" + dataset + ".pq.gt")
	if err != nil {
		fmt.Println("Error creating results file:", err)
		return
	}
	defer outfile.Close()

	// Write header to csv file
	fmt.Fprintf(outfile, "%d,%d\n", nq, nq)

	//use timer from library
	start := time.Now()

	// For each hyperplane create a priority queue and add all points to it with their distance to that hyperplane
	for i := 0; i < 100; i++ {
		fmt.Fprintf(outfile, "%d", i+1)

		// Create priority queue
		pq := make(priorityqueue.DistancePriorityQueue, 0)
		heap.Init(&pq)

		// Add all points to priority queue
		for j := range points {
			dist := hyperplanes[i].dist(&points[j])
			item := &priorityqueue.PQPointDist2Q{
				Point: points[j],
				Dist:  dist,
			}
			heap.Push(&pq, item)
		}

		// Get 100 nearest neighbors
		for j := 0; j < 100; j++ {
			item := heap.Pop(&pq).(*priorityqueue.PQPointDist2Q)
			fmt.Fprintf(outfile, ",%d,%.9f", item.Point.ID+1, item.Dist)
		}
		fmt.Fprintln(outfile)
	}
	elapsed := time.Since(start)
	fmt.Printf("PQ took %s\n", elapsed)
}

func (h *Hyperplane) dist(p *priorityqueue.Point) float32 {
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
