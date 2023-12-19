package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Helper function to check for errors
//
//	If an error is found, the program will panic
//
// Parameters:
// e - The error to check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	ByteCount = iota
	LineCount
	WordCount
	CharCount
)

// Returns a bufio.SplitFunc based on the countType
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

// Returns the count of the specified type
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

type Flags struct {
	FileName  string
	ByteCount bool
	LineCount bool
	WordCount bool
	CharCount bool
}

// Returns a Flags object based on the flags passed in by the user
//
// Returns:
// A Flags object
func getFlags() *Flags {
	// Define flags
	fileNamePtr := flag.String("f", "", "file path")
	byteCountPtr := flag.Bool("c", false, "print the byte counts")
	lineCountPtr := flag.Bool("l", false, "print the newline counts")
	wordCountPtr := flag.Bool("w", false, "print the word counts")
	charCountPtr := flag.Bool("m", false, "print the character counts")

	// Parse flags
	flag.Parse()

	// Create Flags object
	flags := &Flags{
		FileName:  *fileNamePtr,
		ByteCount: *byteCountPtr,
		LineCount: *lineCountPtr,
		WordCount: *wordCountPtr,
		CharCount: *charCountPtr,
	}

	return flags
}

// Prints the counts based on the flags passed in by the user
//
// If no flags are passed in, all counts will be printed
func ccwc() {
	flags := getFlags()

	// Iterate through flags and print counts
	if flags.ByteCount {
		byteCount := fileCounter(flags.FileName, ByteCount)
		fmt.Println("Byte Count:", byteCount)
	}

	if flags.LineCount {
		lineCount := fileCounter(flags.FileName, LineCount)
		fmt.Println("Line Count:", lineCount)
	}

	if flags.WordCount {
		wordCount := fileCounter(flags.FileName, WordCount)
		fmt.Println("Word Count:", wordCount)
	}

	if flags.CharCount {
		charCount := fileCounter(flags.FileName, CharCount)
		fmt.Println("Character Count:", charCount)
	}

	// If no flags are passed in, print all counts
	if !flags.ByteCount && !flags.LineCount && !flags.WordCount && !flags.CharCount {
		byteCount := fileCounter(flags.FileName, ByteCount)
		lineCount := fileCounter(flags.FileName, LineCount)
		wordCount := fileCounter(flags.FileName, WordCount)
		charCount := fileCounter(flags.FileName, CharCount)

		fmt.Println("Byte Count:", byteCount)
		fmt.Println("Line Count:", lineCount)
		fmt.Println("Word Count:", wordCount)
		fmt.Println("Character Count:", charCount)
	}
}

func main() {
	ccwc()
}
