package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename := os.Args[1]

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stdout, "Failed to open file...")
		return
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		i, _ := strconv.Atoi(line)
		sum += i
	}
	fmt.Println(sum)
}
