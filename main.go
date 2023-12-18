// Welcome to the BeeSolver! This nifty package is your secret weapon for tackling
// the New York Times Spelling Bee puzzles. It's not just about finding words; it's
// about embarking on a lexical adventure. With functions to sift through word lists
// and conjure up solutions, BeeSolver transforms a jumble of letters into a celebration
// of words. Whether you're a puzzle enthusiast or a word wizard in the making,
// this is where your journey with letters begins. Let's dive into the hive!
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	requiredLetter := getValidInput("Enter the required letter (one letter only): ", 1, "")
	otherLetters := getValidInput("Enter the other letters (six letters, no duplicates): ", 6, requiredLetter)

	fmt.Println("Required Letter:", requiredLetter)
	fmt.Println("Other Letters:", otherLetters)

	// Load words from the file
	words, err := loadWords("assets/possible_answers.txt")
	if err != nil {
		fmt.Println("Error loading words:", err)
		return
	}

	// Generate a map from the list of words where each key is a word from the list and its corresponding value
	// is a string of the unique letters of that word, sorted alphabetically. This map is used for efficiently
	// checking whether a given combination of letters forms a valid word in the filtered word list.
	wordMap := generateMap(words)

	// Find valid words using the input letters. This function filters the words from the wordMap to find those
	// that include the required letter and are composed exclusively of the combined set of the required and
	// other optional letters. The resulting list, validWords, contains all the words that meet the New York Times
	// Spelling Bee puzzle criteria and is sorted alphabetically.
	validWords := findValidWords(requiredLetter, otherLetters, wordMap)

	for _, word := range validWords {
		fmt.Println(word)
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

// getValidInput prompts the user for input based on the provided prompt message. It validates the input to ensure
// it adheres to the rules of the Spelling Bee game: the input must only contain English letters (A-Z),
// must meet the specified length criteria (one letter for the required letter, six for the other letters), and
// must not include any characters specified in the 'exclude' parameter (to prevent duplication of the required letter
// in the list of other letters). The function repeatedly prompts the user until valid input is provided, enforcing
// the game's input rules for letter selection.
func getValidInput(prompt string, length int, exclude string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.ToUpper(strings.TrimSpace(input))

		if isValidInput(input, length, exclude) {
			return input
		}

		fmt.Println("Invalid input. Please try again.")
	}
}

// isValidInput checks if the input is valid based on length, characters, and exclusion criteria.
func isValidInput(input string, length int, exclude string) bool {
	if len(input) != length {
		return false
	}

	charMap := make(map[rune]bool)
	for _, char := range input {
		if !unicode.IsLetter(char) || strings.ContainsRune(exclude, char) {
			return false
		}
		if charMap[char] {
			return false // Duplicate character found
		}
		charMap[char] = true
	}

	return true
}

// findValidWords finds and returns words that contain the required letter and are made only of the allowed letters.
func findValidWords(requiredLetter, otherLetters string, wordMap map[string]string) []string {
	var validWords []string
	allowedLetters := requiredLetter + otherLetters

	for word, sortedUniqueLetters := range wordMap {
		if contains(sortedUniqueLetters, requiredLetter) && onlyContains(sortedUniqueLetters, allowedLetters) {
			validWords = append(validWords, word)
		}
	}

	sort.Strings(validWords) // I like my word lists alphabetical.
	return validWords
}

// contains checks if the string 's' contains the substring 'sub'.
// 'sub' is the letter that Spelling Bee requires for a word to be a valid answer.
func contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

// onlyContains checks if the string 's' contains only the characters in 'allowed'.
// words are only valid answers in Spelling Bee if they only use a combination of
// the required and optional letters.
func onlyContains(s, allowed string) bool {
	for _, char := range s {
		if !strings.ContainsRune(allowed, char) {
			return false
		}
	}
	return true
}
