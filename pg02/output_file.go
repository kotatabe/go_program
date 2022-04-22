package main

import (
	"fmt"
	"os"
	"io"
)

func OutputFile(filename string) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File Open failed")
		return
	}
	defer fp.Close()

	_, err = io.Copy(os.Stdout, fp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File Copy failed")
		return
	}
}
