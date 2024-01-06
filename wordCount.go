package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getWordCount(f string) int {
	ofile := openFile(f)
	defer ofile.Close()
	wc := 0
	scanner := bufio.NewScanner(ofile)
	for scanner.Scan() {
		line := scanner.Text()
		words := getWords(line)
        numOfWords := len(words)
		if *debug {
			fmt.Println(words, " -> ", numOfWords)
			fmt.Println("--------")
		}
		wc += numOfWords
	}
	return wc
}

func getWords(l string) []string {
	words := strings.Fields(l) // this is better than splitting the string on " "
	return words
}
