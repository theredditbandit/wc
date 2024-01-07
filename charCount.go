package main

import (
	"bytes"
	"os"
)

func getCharCount(f string) int {
    file , _ := os.ReadFile(f)
    chars := len(bytes.Runes(file))
    return chars
}
