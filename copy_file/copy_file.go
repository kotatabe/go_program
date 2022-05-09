package copy

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/xerrors"
)

func CopyFile(filename string) {
	src_fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%+v\n", xerrors.New("failed to open file"))
		return
	}

	copy_filename := "copy.txt"
	dst_fp, err := os.Create(copy_filename)
	if err != nil {
		fmt.Printf("%+v\n", xerrors.New("failed to create new file"))
		return
	}

	mw := io.MultiWriter(os.Stdout, dst_fp)

	_, err = io.Copy(mw, src_fp)
	if err != nil {
		fmt.Printf("%+v\n", xerrors.New("failed to copy file"))
		return
	}
}
