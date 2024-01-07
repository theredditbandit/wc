package main

import (
	"fmt"
	"os"
)

func openFile(f string) *os.File {
	file, err := os.Open(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	return file
}

func givenFileIsADir(f string) bool {
	file := openFile(f)
	defer file.Close()
	stat, _ := file.Stat()
	return stat.IsDir()
}

// TODO : Pretty print this with some charmcli goodness ?
func output(bc, lc, cc, wc, toPrint int, fname string) {
	if toPrint == 1 {
		size := bc + lc + cc + wc // since only one value to print rest will be 0
		fmt.Printf("%d %s\n", size, fname)
	} else if *verbose {
		fmt.Printf("line count :%d\nchar count :%d\nbyte count :%d\nword count: %d\n  \t%s\n", lc, cc, bc, wc, fname)
	} else {
		fmt.Printf("  %d  %d  %d %s\n", lc, wc, bc, fname)
	}
}
