package output

import (
	"io"
	"os"

	"golang.org/x/xerrors"
)

func OutputFile(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		return xerrors.New("file open failed")
	}
	defer fp.Close()

	_, err = io.Copy(os.Stdout, fp)
	if err != nil {
		return xerrors.New("file copy failed")
	}
	return nil
}
