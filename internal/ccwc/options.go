// internal/ccwc/options.go

package ccwc

import (
	"flag"
)

type options struct {
	FileName  string
	ByteCount bool
	LineCount bool
	WordCount bool
	CharCount bool
}

// getOptions returns an options object based on the flags passed in by the user
//
// Returns:
// An options object
func getOptions() *options {
	// Define flags
	byteCountPtr := flag.Bool("c", false, "print the byte counts")
	lineCountPtr := flag.Bool("l", false, "print the newline counts")
	wordCountPtr := flag.Bool("w", false, "print the word counts")
	charCountPtr := flag.Bool("m", false, "print the character counts")

	// Parse flags
	flag.Parse()

	// Parse arguments
	args := flag.Args()

	// Get file name
	fileName := ""
	if len(args) != 0 {
		fileName = args[0]
	}

	// Create options object
	options := &options{
		FileName:  fileName,
		ByteCount: *byteCountPtr,
		LineCount: *lineCountPtr,
		WordCount: *wordCountPtr,
		CharCount: *charCountPtr,
	}

	return options
}
