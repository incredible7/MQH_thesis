package utils

import (
    "encoding/binary"
    "fmt"
    "math"
    "os"
    "MQH_THESIS/pkg/types"
)

// ReadBinaryFile reads a binary file and returns its contents
func ReadBinaryFile(filepath string) ([]byte, error) {
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

// ReadPoints reads points from binary data
func ReadPoints(data []byte, n int, d int) []types.Point {
	totalValues := d + 1 // Each entry has d float32 + 1 delimiter - the delimiter is used
	points := make([]types.Point, n)

	for i := 0; i < n; i++ {
		startIndex := i * totalValues

		// Read feature vector (first d values)
		points[i] = types.Point{
			ID:          i, // Assign sequential ID
			Coordinates: make([]float32, d),
		}
		for j := 0; j < d; j++ {
			bits := binary.LittleEndian.Uint32(data[(startIndex+j)*4 : (startIndex+j+1)*4])
			points[i].Coordinates[j] = math.Float32frombits(bits)
		}
		// Ignore the last value (which is always 1.0)

	}

	fmt.Printf("Loaded points 🧙 \n")
	return points
}

// ReadHyperplanes reads hyperplanes from binary data
func ReadHyperplanes(querydata []byte, d int) []types.Hyperplane {
	hyperplanes := make([]types.Hyperplane, 100)
	for i := 0; i < 100; i++ {
		query := types.Hyperplane{
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
	fmt.Printf("Loaded hyperplanes 🧹 \n")
	return hyperplanes
}