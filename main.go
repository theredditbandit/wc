package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	byteCount = flag.Bool("c", false, "print the byte count")
	charCount = flag.Bool("m", false, "print the char count")
	lineCount = flag.Bool("l", false, "print the line count")
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 { // TODO : Iterate through all the files passed instead of just processing the first one, should be easy
		flag.Usage()
		os.Exit(1)
	}
	fname := flag.Arg(0)

	if !fileExists(fname) {
		fmt.Printf("wc: %s: No such file or directory\n", fname)
		os.Exit(1)
	}
	if givenFileIsADir(fname) {
		fmt.Printf("wc: %s: Is a directory\n", fname)
		os.Exit(1)
	}

	var bc, cc, lc int
	if *byteCount {
		bc = getByteCount(fname)
	}
	if *charCount {
		cc = getCharCount(fname)
	}
	if *lineCount {
		lc = getLineCount(fname)
	}
	if !*byteCount && !*charCount && !*lineCount {
		bc, cc, lc = getByteCount(fname), getCharCount(fname), getLineCount(fname)
	}

	output(bc, cc, lc, fname)

}
