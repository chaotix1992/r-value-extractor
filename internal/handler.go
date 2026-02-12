package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ReadCSV reads a CSV file from the given path and returns [][]string.
func ReadCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv: %w", err)
	}

	return records, nil
}

// WriteCSV writes [][]string data to a CSV file at the given path.
func WriteCSV(data [][]string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		return fmt.Errorf("failed to write csv: %w", err)
	}

	return nil
}

func FilterCSV(p_values [][]string, r_values [][]string, max_val float64) [][]string {
	// Iterate over first array. If value > 0.03, discard the respective r_value entry, else keep it
	filtered := [][]string{}

	for i := range(len(p_values)) {
		row := []string{}
		for j := range(len(p_values[i])) {
			p_value, err := strconv.ParseFloat(p_values[i][j], 64)
			if err != nil {
				row = append(row, r_values[i][j])
			} else {
				if p_value > max_val {
					row = append(row, "")
				} else {
					row = append(row, r_values[i][j])
				}
			}
		}
		filtered = append(filtered, row)
	}
	return filtered
}
