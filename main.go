package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <base64>")
		fmt.Println("Example: go run main.go SGVsbG8gV29ybGQ=")
		os.Exit(1)
	}

	base64Input := os.Args[1]

	// Generate filename with current timestamp
	outputFile := generateFilename()

	// Decode base64 input
	decodedData, err := base64.StdEncoding.DecodeString(base64Input)
	if err != nil {
		fmt.Println("Error decoding base64 input:", err)
		fmt.Println("Please ensure your base64 input is valid.")
		os.Exit(1)
	}

	// Write the decoded data to the output file
	err = os.WriteFile(outputFile, decodedData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}

	fmt.Printf("File saved as %s\n", outputFile)
}

func generateFilename() string {
	// Get current timestamp
	now := time.Now()
	// Format the timestamp as a string (YYYYMMDD_HHMMSS)
	timestamp := now.Format("20060102_150405")
	// Create the filename
	return fmt.Sprintf("%s.xlsx", timestamp)
}
