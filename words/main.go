package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("filtered_words.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	// Process each word in the input file
	for scanner.Scan() {
		word := scanner.Text()
		// Check the length of the word
		if len(word) >= 4 && len(word) <= 10 {
			_, err := writer.WriteString(word + "\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Flush the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing output file:", err)
		return
	}

	fmt.Println("File processed successfully.")
}
