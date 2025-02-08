package main

import (
	"MQH_THESIS/pkg/algorithms"
	"MQH_THESIS/pkg/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Error - try again using this format:\n go run . <dataset name> <num_points> <dimensionality> <num queries> <k>")
		return
	}

	// Read command-line arguments
	dataset := os.Args[1]
	n, _ := strconv.Atoi(os.Args[2])
	d, _ := strconv.Atoi(os.Args[3])
	nq, _ := strconv.Atoi(os.Args[4])
	k, _ := strconv.Atoi(os.Args[5])

	// Load data
	pointsData, err := utils.ReadBinaryFile("data/datasets/" + dataset + ".ds")
	if err != nil {
		fmt.Printf("Error reading dataset file: %v\n", err)
		return
	}

	queriesData, err := utils.ReadBinaryFile("data/datasets/" + dataset + ".q")
	if err != nil {
		fmt.Printf("Error reading query file: %v\n", err)
		return
	}

	points := utils.ReadPoints(pointsData, n, d)
	hyperplanes := utils.ReadHyperplanes(queriesData, d)

	// Determine mode and suffix
	isGroundTruth := k == 100 && nq == 100
	suffix := ".gt"
	if !isGroundTruth {
		suffix = fmt.Sprintf(".nq%d.k%d", nq, k)
	}

	// Run algorithms
	// algorithms.ExhaustiveFS(dataset, points, hyperplanes, nq, k, suffix)
	// algorithms.ExhaustivePQ(dataset, points, hyperplanes, nq, k, suffix)
	levels := 3
	algorithms.Mqh(dataset, points, hyperplanes, nq, d, n, suffix, levels)
	

	// Add your new algorithm here
	// algorithms.NewAlgorithm(dataset, points, hyperplanes, nq, k, suffix)
}
