package main

import (
	"fmt"
	"os"
	"github.com/kotatabe/go_program/pg03/sumfile"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename := os.Args[1]
	sum, err := sumfile.SumFile(filename)
	if err != nil {
		fmt.Errorf("Error: %w", err)
		return
	}
	fmt.Println(sum)
}
