# Word Count Clone

## Description

An interpretation of the Unix command line tool `wc`.

Solution to John' Crickett's [Coding Challenge #1 Build your own wc](https://codingchallenges.substack.com/p/coding-challenge-1) written in Golang.

## Pre-requisites

Requires you to have [golang](https://go.dev/doc/install) installed

The Makefile includes a `lint` command which depends on [golangci-lint](https://github.com/golangci/golangci-lint). If you want to use `make lint`, then install this dependency.

## Installation

1. Clone and navigate into the repository
2. Build the executable binaries using `make build`
3. The binaries can be found in the `bin` directory of the repo

## Usage

Reading counts from passed in file
```bash
# Byte Count
./ccwc -c <filepath>

# Line Count
./ccwc -l <filepath>

# Word Count
./ccwc -w <filepath>

# Char Count
./ccwc -m <filepath>

# All Counts (default w/ no options)
./ccwc <filepath>
```

Reading counts from stdin
```bash
# Byte Count
cat <filepath> | ./ccwc -c

# Line Count
cat <filepath> | ./ccwc -l

# Word Count
cat <filepath> | ./ccwc -w

# Char Count
cat <filepath> | ./ccwc -m

# All Counts (default w/ no options)
cat <filepath> | ./ccwc <filepath>
```

Example Usage
```bash
./bin/ccwc ../tests/test.txt
Byte Count: 335040
Line Count: 7143
Word Count: 58164
Character Count: 332144

cat /tests/test.txt | ./bin/ccwc
Byte Count: 335040
Line Count: 7143
Word Count: 58164
Character Count: 332144
```

## Resources / References
- [John Crickett - Coding Challenge #1 - Build your own wc!](https://codingchallenges.substack.com/p/coding-challenge-1)
- [Melkey - This is the BEST Way to Structure Your GO Projects](https://www.youtube.com/watch?v=dxPakeBsgl4)
- [Anthony GG - How I structure New Projects in Golang](https://www.youtube.com/watch?v=dJIUxvfSg6A)
- [wc man page](https://linux.die.net/man/1/wc)
- [Alex Edwards - Quick tip: A time-saving Makefile for your Go projects](https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects)
- [Air - Live reload for Go apps](https://github.com/cosmtrek/air)
- [Go By Example - Command Line Arguments](https://gobyexample.com/command-line-arguments)
- [Go By Example - Command Line Flags](https://gobyexample.com/command-line-flags)
- [Go By Example - Reading Files](https://gobyexample.com/reading-files)
- [Go By Example - Temporary Files & Directories](https://gobyexample.com/temporary-files-and-directories)
- [NerdCademy - Learn to Read Files w/ Go!](https://www.youtube.com/watch?v=8uKtaYJlTzs)
- [Honeybadger - A comprehensive guide to file operations in Go](https://www.honeybadger.io/blog/comprehensive-guide-to-file-operations-in-go/)
- [spf13/cobra](https://github.com/spf13/cobra)
- [Melkey - This makes Golang CLIL Development So MUCH Better](https://www.youtube.com/watch?v=yybzcix10XI)
- [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)
- [Melkey - This is Why You NEED to Use This Golang CLI Framework](https://www.youtube.com/watch?v=ncakAFWxIys)
- [Jetbrains - Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)