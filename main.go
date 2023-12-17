package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// Load words from the file
	words, err := loadWords("words/filtered_words.txt")
	if err != nil {
		fmt.Println("Error loading words:", err)
		return
	}

	// Generate the map
	wordMap := generateMap(words)

	// Printing a small portion of the map for demonstration
	for k, v := range wordMap {
		fmt.Println(k, ":", v)
		// Example: Break after printing 10 items for demonstration
		break
	}
}

// loadWords loads words from a given file.
func loadWords(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

// generateMap generates a map of words to their sorted unique letters.
func generateMap(words []string) map[string]string {
	wordMap := make(map[string]string)
	for _, word := range words {
		wordMap[word] = sortUniqueLetters(word)
	}
	return wordMap
}

// sortUniqueLetters returns a string with alphabetically sorted unique letters of the given word.
func sortUniqueLetters(word string) string {
	letterMap := make(map[rune]bool)
	for _, letter := range word {
		letterMap[letter] = true
	}

	var uniqueLetters []rune
	for letter := range letterMap {
		uniqueLetters = append(uniqueLetters, letter)
	}

	sort.Slice(uniqueLetters, func(i, j int) bool {
		return uniqueLetters[i] < uniqueLetters[j]
	})

	return string(uniqueLetters)
}
