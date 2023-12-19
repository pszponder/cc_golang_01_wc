// internal/ccwc/flags.go

package ccwc

import (
	"flag"
)

type Flags struct {
	FileName  string
	ByteCount bool
	LineCount bool
	WordCount bool
	CharCount bool
}

// getFlags returns a Flags object based on the flags passed in by the user
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
