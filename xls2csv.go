package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func cleanText(text string) string {
	// Replace any combination of \r\n or \n with a space
	return strings.ReplaceAll(strings.ReplaceAll(text, "\r\n", " "), "\n", " ")
}

func excelToCSV(filePath string) error {
	// Open Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("error opening Excel file: %w", err)
	}
	defer f.Close()

	// Get all sheet names
	sheetNames := f.GetSheetList()

	// Process each sheet
	for _, sheetName := range sheetNames {
		// Read all rows from the sheet
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return fmt.Errorf("error reading sheet %s: %w", sheetName, err)
		}

		// Clean data: replace line breaks with spaces
		cleanRows := make([][]string, len(rows))
		for i, row := range rows {
			cleanRows[i] = make([]string, len(row))
			for j, cell := range row {
				cleanRows[i][j] = cleanText(cell)
			}
		}

		// Create CSV file
		csvFileName := sheetName + ".csv"
		csvFile, err := os.Create(csvFileName)
		if err != nil {
			return fmt.Errorf("error creating CSV file %s: %w", csvFileName, err)
		}
		defer csvFile.Close()

		// Create CSV writer
		writer := csv.NewWriter(csvFile)
		defer writer.Flush()

		// Write all rows to CSV
		for _, row := range cleanRows {
			if err := writer.Write(row); err != nil {
				return fmt.Errorf("error writing to CSV file %s: %w", csvFileName, err)
			}
		}

		fmt.Printf("Worksheet '%s' has been exported to '%s'\n", sheetName, csvFileName)
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <excel_file>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: File '%s' not found\n", filePath)
		os.Exit(1)
	}

	if err := excelToCSV(filePath); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
