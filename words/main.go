// This is a little helper script to slightly reduce the words list based on the game rules
// It really only reduces the possible answer set by ~25%, but I kept this around since I already wrote it
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file, which is just a scrabble dictionary.
	// The New York Times does not maintain an official word list.
	inputFile, err := os.Open("../assets/words.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("../assets/possible_answers.txt")
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
