package sumfile

import (
	"bufio"
	"os"
	"strconv"

	"golang.org/x/xerrors"
)

func SumFile(filename string) (int, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return -1, xerrors.New("failed to open file")
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			return -1, xerrors.New("failed to open file")
		}
		sum += i
	}
	return sum, nil
}
