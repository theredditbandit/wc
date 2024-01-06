package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getWordCount(f string) int {
	ofile := openFile(f)
	defer ofile.Close()
	lc := 0
	wc := 0
	scanner := bufio.NewScanner(ofile)
	for scanner.Scan() {
		lc++
		line := scanner.Text()
		words := countWords(line)
		if *debug {
			fmt.Printf("words: %v\n", words)
			fmt.Println("--------")
		}
		wc += words
	}
	return wc
}

func countWords(l string) int {
    words := strings.Fields(l) // this is better than splitting the string on " "
	if len(words) != 0 && *debug {
		fmt.Println(words, " -> ", len(words))
	}
	return len(words)
}
