package main

import (
	"fmt"
	"os"
)

func openFile(f string) *os.File {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Something went wrong: ", err)
	}
	return file
}

func fileExists(f string) bool {

	return true
}

func givenFileIsADir(f string) bool {
	file := openFile(f)
	defer file.Close()
	stat, _ := file.Stat()
	return stat.IsDir()
}

func output(bc, lc, wc int, name string) {
	fmt.Printf("%d %d %d %s\n", lc, wc, bc, name)
}
