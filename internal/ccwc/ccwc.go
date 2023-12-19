// internal/ccwc/ccwc.go

package ccwc

import "fmt"

// Ccwc prints the counts based on the flags passed in by the user
//
// If no flags are passed in, all counts will be printed
func Ccwc() {
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
