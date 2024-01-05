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

func output(bc, lc, wc, toPrint int, name string) {
	if toPrint == 1 {
		size := bc + lc + wc // since only one value to print rest will be 0
		fmt.Printf("%d %s\n", size, name)
	} else {
		fmt.Printf("%d %d %d %s\n", lc, wc, bc, name)
	}
}
