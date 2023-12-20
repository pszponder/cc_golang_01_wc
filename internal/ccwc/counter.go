// internal/ccwc/counter.go

package ccwc

import (
	"bufio"
	"os"
)

const (
	ByteCount = iota
	LineCount
	WordCount
	CharCount
)

// scanType returns a bufio.SplitFunc based on the countType
//
//	The bufio.SplitFunc is used to split the input into tokens when scanning
//
// Parameters:
// scanType - The type of scan to use
//
//	ByteCount: Number of bytes
//	LineCount: Number of lines
//	WordCount: Number of words
//	CharCount: Number of characters
//
// Returns:
// A bufio.SplitFunc
func scanType(scanType int) bufio.SplitFunc {
	switch scanType {
	case ByteCount:
		return bufio.ScanBytes
	case LineCount:
		return bufio.ScanLines
	case WordCount:
		return bufio.ScanWords
	case CharCount:
		return bufio.ScanRunes
	default:
		panic("Invalid scanType")
	}
}

// fileCounter returns the count of the specified type
//
// Parameters:
// fileName - The name of the file to count
// countType - The type of count to return
//
//	Count Types:
//	ByteCount: Number of bytes
//	LineCount: Number of lines
//	WordCount: Number of words
//	CharCount: Number of characters
//
// Returns:
// The count of the specified type
func fileCounter(fileName string, countType int) int {
	// Open File
	file, err := os.Open(fileName)
	check(err)
	defer file.Close() // Close file when function returns

	// Count
	scanner := bufio.NewScanner(file)
	scanner.Split(scanType(countType))
	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}