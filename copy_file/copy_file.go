package copy

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(filename string) {
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

	mw := io.MultiWriter(os.Stdout, dst_fp)

	_, err = io.Copy(mw, src_fp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to copy file...")
		return
	}
}
