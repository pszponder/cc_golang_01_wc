package main

import (
	"fmt"

	"github.com/pszponder/cc_golang_01_wc/internal/ccwc"
)

func main() {
	opts := ccwc.ParseOptions()

	counts, err := ccwc.Counter(opts.FileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate through options and print counts
	if opts.ByteCount {
		fmt.Println("Byte Count:", counts.ByteCount)
	}

	if opts.LineCount {
		fmt.Println("Line Count:", counts.LineCount)
	}

	if opts.WordCount {
		fmt.Println("Word Count:", counts.WordCount)
	}

	if opts.CharCount {
		fmt.Println("Character Count:", counts.CharCount)
	}

	// If no options are passed in, print all counts
	if !opts.ByteCount && !opts.LineCount && !opts.WordCount && !opts.CharCount {
		fmt.Println("Byte Count:", counts.ByteCount)
		fmt.Println("Line Count:", counts.LineCount)
		fmt.Println("Word Count:", counts.WordCount)
		fmt.Println("Character Count:", counts.CharCount)
	}
}
