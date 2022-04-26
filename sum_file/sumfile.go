package sumfile

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func SumFile(filename string) (int, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return -1, errors.New("failed to open file")
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			return -1, errors.New("failed to open file")
		}
		sum += i
	}
	return sum, nil
}
