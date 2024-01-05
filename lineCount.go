package main

import "bufio"

func getLineCount(f string) int {
	ofile := openFile(f)
	defer ofile.Close()
    lc := 0 
    scanner := bufio.NewScanner(ofile)
    for scanner.Scan() {
        lc ++
    }
    return lc 
}
