package parsechunk

import (
	"encoding/csv"
	"os"
)

// ReadAndChunkCSV reads a CSV file and organizes the data by columns.
func ReadAndChunkCSV(filePath string) (map[string][]string, error) {
	// Opening the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// CSV reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Organizing data by columns
	columns := make(map[string][]string)
	for rowIndex, row := range records {
		for columnIndex, cell := range row {
			if rowIndex == 0 { // Assuming the first row contains column headers
				columns[cell] = make([]string, 0)
			} else {
				header := records[0][columnIndex]
				columns[header] = append(columns[header], cell)
			}
		}
	}

	return columns, nil
}

// CreateAndWriteCSV creates a CSV file for each column with its respective data.
// The function now includes the output directory parameter for flexibility.
func CreateAndWriteCSV(columnName string, data []string, outputDir string) error {
	filePath := outputDir + columnName + ".csv"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(data); err != nil { // Data already includes header
		return err
	}

	return nil
}
