// internal/ccwc/ccwc.go

package ccwc

import "fmt"

// Ccwc prints the counts based on the options passed in by the user
//
// If no flags are passed in, all counts will be printed
// If no file name is passed in, stdin will be scanned instead
func Ccwc() {
	options := getOptions()

	// Iterate through options and print counts
	if options.ByteCount {
		byteCount := counter(options.FileName, ByteCount)
		fmt.Println("Byte Count:", byteCount)
	}

	if options.LineCount {
		lineCount := counter(options.FileName, LineCount)
		fmt.Println("Line Count:", lineCount)
	}

	if options.WordCount {
		wordCount := counter(options.FileName, WordCount)
		fmt.Println("Word Count:", wordCount)
	}

	if options.CharCount {
		charCount := counter(options.FileName, CharCount)
		fmt.Println("Character Count:", charCount)
	}

	// If no options are passed in, print all counts
	if !options.ByteCount && !options.LineCount && !options.WordCount && !options.CharCount {
		byteCount := counter(options.FileName, ByteCount)
		lineCount := counter(options.FileName, LineCount)
		wordCount := counter(options.FileName, WordCount)
		charCount := counter(options.FileName, CharCount)

		fmt.Println("Byte Count:", byteCount)
		fmt.Println("Line Count:", lineCount)
		fmt.Println("Word Count:", wordCount)
		fmt.Println("Character Count:", charCount)
	}
}
