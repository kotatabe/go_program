package sumfile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/xerrors"
)

func SumFile(filename string) (int, error) {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%+v\n", xerrors.New("failed to open file"))
		return -1, err
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%+v\n", xerrors.New("failed to open file"))
			return -1, err
		}
		sum += i
	}
	return sum, nil
}
