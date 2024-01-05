package main

func getByteCount(f string) int {
	ofile := openFile(f)
	defer ofile.Close()
	stat, _ := ofile.Stat()
	return int(stat.Size())
}
