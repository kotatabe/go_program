package copy

import (
	"io"
	"os"

	"golang.org/x/xerrors"
)

func CopyFile(filename string) error {
	src_fp, err := os.Open(filename)
	if err != nil {
		return xerrors.New("failed to open file")
	}

	copy_filename := "copy.txt"
	dst_fp, err := os.Create(copy_filename)
	if err != nil {
		return xerrors.New("failed to create new file")
	}

	mw := io.MultiWriter(os.Stdout, dst_fp)

	_, err = io.Copy(mw, src_fp)
	if err != nil {
		return xerrors.New("failed to copy file")
	}
	return nil
}
