package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strconv"
)

var dist_map = make(map[int]float32)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <dataset name> <num_points> <dimensionality>")
		return
	}

	// Read command-line arguments
	dataset := os.Args[1]
	n, _ := strconv.Atoi(os.Args[2])
	d, _ := strconv.Atoi(os.Args[3])
	k, _ := strconv.Atoi(os.Args[4])

	fmt.Printf("🔍 Loading dataset: %s with %d points of %d dimensions\n", dataset, n, d)

	// Open the binary data file
	datafile, err := os.Open("data/datasets/" + dataset + ".ds")
	if err != nil {
		fmt.Println("Error opening datafile:", err)
		return
	}
	defer datafile.Close()

	// Open the binary query file
	queryfile, err := os.Open("data/datasets/" + dataset + ".q")
	if err != nil {
		fmt.Println("Error opening query file:", err)
		return
	}
	defer queryfile.Close()

	// Read data file
	datastat, _ := datafile.Stat()
	data := make([]byte, datastat.Size())
	_, err = datafile.Read(data)
	if err != nil {
		fmt.Println("Error reading data file:", err)
		return
	}

	// read query file
	querystat, _ := queryfile.Stat()
	fmt.Printf("size of query input : %d\n", querystat.Size())
	querydata := make([]byte, querystat.Size())
	_, err = queryfile.Read(querydata)
	if err != nil {
		fmt.Println("Error reading query file:", err)
		return
	}

	totalValues := d + 1 // Each entry has d float32 + 1 delimiter
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		startIndex := i * totalValues

		// Read feature vector (first d values)
		points[i] = Point{
			ID:          i, // Assign sequential ID
			Coordinates: make([]float32, d),
		}
		for j := 0; j < d; j++ {
			bits := binary.LittleEndian.Uint32(data[(startIndex+j)*4 : (startIndex+j+1)*4])
			points[i].Coordinates[j] = math.Float32frombits(bits)
		}

		// Ignore the last value (which is always 1.0)
	}

	fmt.Printf("✅ Loaded %d points\n", n)

	// check that points have
	for i := 0; i < n; i++ {
		if len(points[i].Coordinates) != d {
			fmt.Printf("point number %d doesn't have d = %d\n", i, d)
		}
	}

	// read query
	hyperplanes := make([]Hyperplane, 100)
	for i := 0; i < 100; i++ {
		query := Hyperplane{
			Q: make([]float32, d),
			B: 0.0,
		}
		for i := 0; i < d; i++ {
			bits := binary.LittleEndian.Uint32(querydata[i*4 : (i+1)*4])
			query.Q[i] = math.Float32frombits(bits)
		}
		Bbits := binary.LittleEndian.Uint32(querydata[d*4 : (d+1)*4])
		query.B = math.Float32frombits(Bbits)
		hyperplanes[i] = query
	}

	for x := range len(hyperplanes) {
		fmt.Printf("first 5 values of hyperplane %d's normal vector : %v\n", x, hyperplanes[x].Q[0:5])
		fmt.Printf("bias of hyperplane %d: %v\n", x, hyperplanes[x].B)
	}
	nns := make([][]Point, 100)
	for i := 0; i < 100; i++ {
		nns[i] = make([]Point, k)
		current := float32(math.SmallestNonzeroFloat32)
		for j := 0; j < k; j++ {
			dist := hyperplanes[i].dist(&points[j])
			dist_map[points[j].ID] = dist
			if dist > current {
				current = dist
			}
			nns[i][j] = points[j]
		}
		for j := k; j < len(points); j++ {
			dist := hyperplanes[i].dist(&points[j])
			if dist < current {
				current = dist
				Add(&nns[i], &points[j], k)
			}
		}
	}

}

func Add(l *[]Point, p *Point, k int) {
	for i := 0; i < k; i++ {
		if dist_map[l[i].ID] > dist_map[p.ID] {

		}
	}

}

type Point struct {
	ID          int
	Coordinates []float32
}

type Hyperplane struct {
	Q []float32
	B float32
}

func (h *Hyperplane) dist(p *Point) float32 {
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
	denominator = float32(math.Sqrt(float64(denominator)))
	return numerator / denominator
}
