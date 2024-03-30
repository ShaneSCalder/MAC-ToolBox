package main

import (
	"CHUNK/pkg/parsechunk" // Update this path according to your project structure
	"fmt"
	"math"
)

func main() {
	inputCSV := "data/input.csv"
	outputDir := "dataout/"
	portion := 0.25 // Since you want 25% of the data in each file

	columns, err := parsechunk.ReadAndChunkCSV(inputCSV)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for columnName, data := range columns {
		totalRows := len(data) - 1 // Excluding header
		rowsPerChunk := int(math.Ceil(float64(totalRows) * portion))
		if rowsPerChunk == 0 {
			rowsPerChunk = 1 // Ensure at least one row per chunk, adjust based on your requirements
		}

		totalChunks := int(math.Ceil(float64(totalRows) / float64(rowsPerChunk)))

		for i := 0; i < totalChunks; i++ {
			start := i*rowsPerChunk + 1 // Skip header for data, +1 to adjust for header
			end := start + rowsPerChunk
			if end > totalRows+1 { // Adjust for header in totalRows
				end = totalRows + 1
			}

			chunk := make([]string, 0, end-start+1)
			chunk = append(chunk, data[0]) // Add header
			chunk = append(chunk, data[start:end]...)

			chunkFileName := fmt.Sprintf("%s_part_%d.csv", columnName, i+1)
			err := parsechunk.CreateAndWriteCSV(chunkFileName, chunk, outputDir)
			if err != nil {
				fmt.Println("Error writing chunk for column", columnName, ":", err)
				continue
			}
			fmt.Println("Successfully created chunk:", chunkFileName)
		}
	}
}
