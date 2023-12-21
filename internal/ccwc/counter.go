// internal/ccwc/counter.go

package ccwc

import (
	"bufio"
	"os"
	"strings"
	"unicode/utf8"
)

// CountResult is a struct that contains the results of a count
type CountResult struct {
	ByteCount int
	LineCount int
	WordCount int
	CharCount int
}

// Counter counts the bytes, lines, words, and characters in a file
//
// Parameters:
// fileName: The name of the file to count (if empty, stdin will be scanned)
//
// Returns:
// A CountResult struct containing the counts
func Counter(fileName string) (CountResult, error) {
	var scanner *bufio.Scanner

	if fileName == "" {
		// Scan stdin
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		// Open File
		file, err := os.Open(fileName)
		if err != nil {
			return CountResult{}, err
		}
		defer file.Close() // Close file when function returns

		// Scan file
		scanner = bufio.NewScanner(file)
	}

	// Initialize counts (0 by default)
	counts := CountResult{}

	// Count (scan line by line)
	for scanner.Scan() {
		line := scanner.Text() // Get line of text from the scanner

		counts.LineCount++                               // Increment line count
		counts.ByteCount += len([]byte(line))            // Convert line to byte slice & get the length
		counts.WordCount += len(strings.Fields(line))    // Convert line to slice of words & get the length
		counts.CharCount += utf8.RuneCountInString(line) // Get the length of the line
	}

	return counts, scanner.Err()
}
