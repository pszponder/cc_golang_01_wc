package main

import (
	"fmt"

	"github.com/pszponder/cc_golang_01_wc/internal/ccwc"
)

func main() {
	options := ccwc.GetOptions()

	counts, err := ccwc.Counter(options.FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through options and print counts
	if options.ByteCount {
		fmt.Println("Byte Count:", counts.ByteCount)
	}

	if options.LineCount {
		fmt.Println("Line Count:", counts.LineCount)
	}

	if options.WordCount {
		fmt.Println("Word Count:", counts.WordCount)
	}

	if options.CharCount {
		fmt.Println("Character Count:", counts.CharCount)
	}

	// If no options are passed in, print all counts
	if !options.ByteCount && !options.LineCount && !options.WordCount && !options.CharCount {
		fmt.Println("Byte Count:", counts.ByteCount)
		fmt.Println("Line Count:", counts.LineCount)
		fmt.Println("Word Count:", counts.WordCount)
		fmt.Println("Character Count:", counts.CharCount)
	}
}
