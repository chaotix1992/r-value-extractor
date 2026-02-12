package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"r-value-extractor/internal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter R-Value CSV file path: ")
	inputPath1, _ := reader.ReadString('\n')

	fmt.Print("Enter P-Value CSV file path: ")
	inputPath2, _ := reader.ReadString('\n')

	fmt.Print("Enter the maximum P-Value value that should be filtered: ")
	maximum, _ := reader.ReadString('\n')

	fmt.Print("Enter output CSV file path: ")
	outputPath, _ := reader.ReadString('\n')

	// Remove newline characters
	inputPath1 = trimNewline(inputPath1)
	inputPath2 = trimNewline(inputPath2)
	outputPath = trimNewline(outputPath)
	maximum = trimNewline(maximum)
	max_val, _ := strconv.ParseFloat(maximum, 64)

	// Read first CSV
	data1, err := internal.ReadCSV(inputPath1)
	if err != nil {
		fmt.Println("Error reading first CSV:", err)
		return
	}

	// Read second CSV
	data2, err := internal.ReadCSV(inputPath2)
	if err != nil {
		fmt.Println("Error reading second CSV:", err)
		return
	}

	// Combine data
	combined := internal.FilterCSV(data2, data1, max_val)

	// Write output CSV
	err = internal.WriteCSV(combined, outputPath)
	if err != nil {
		fmt.Println("Error writing output CSV:", err)
		return
	}

	fmt.Println("CSV files successfully filtered and written to", outputPath)
	fmt.Println("Press ENTER to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Helper function to trim newline characters
func trimNewline(s string) string {
	if len(s) == 0 {
		return s
	}
	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '\r' {
		s = s[:len(s)-1]
	}
	return s
}
