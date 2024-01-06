package main

import (
	"bufio"
	"fmt"
	"strings"
	// "text/scanner"
	// "unicode"
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
		words := countWords(strings.TrimSpace(line))
		if *debug {
			fmt.Printf("words: %v\n", words)
			fmt.Println("--------")
		}
		wc += words
		// fmt.Printf("wc: %v\n", wc)
	}
	return wc
}

func countWords(l string) int {
	words := strings.Split(l, " ")
	// var s scanner.Scanner
	// s.Init(strings.NewReader(l))
	// var wc int
	//
	// for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
	//     if tok == scanner.Ident {
	//         wc ++
	//     }
	// }
	// words := strings.FieldsFunc(l, func(r rune) bool {
	// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	// })
	if len(words) != 0 && *debug {
		fmt.Println(words, " -> ", len(words))
	}
	if words[0] == "" {
		return 0
	}
	// return len(words)

	return len(words)
}
