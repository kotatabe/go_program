package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename := os.Args[1]

	src_fp, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open file...")
		return
	}

	copy_filename := "copy.txt"
	dst_fp, err := os.Create(copy_filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create new file...")
		return
	}

	_, err = io.Copy(dst_fp, src_fp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to copy file...")
		return
	}

	// コピーしたファイルを出力
	OutputFile(copy_filename)
}