// copy file 

package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
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

	dst_fp, err := os.Create("copy.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create new file...")
		return
	}

	_, err = io.Copy(io.Writer(dst_fp), io.Reader(src_fp))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to copy file...")
		return
	}

	// コピーしたファイルを出力
	_, err = io.Copy(os.Stdout, io.Reader(dst_fp))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to copy file...")
		return
	}
}