package output

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/xerrors"
)

func OutputFile(filename string) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%+v\n", xerrors.New("file open failed"))
		return
	}
	defer fp.Close()

	_, err = io.Copy(os.Stdout, fp)
	if err != nil {
		fmt.Printf("%+v\n", xerrors.New("file copy failed"))
		return
	}
}
