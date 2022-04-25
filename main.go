package main

import (
	"fmt"
	// "github.com/kotatabe/go_program/files"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename := os.Args[1]

	// OutputFile(filename)

	// CopyFile(filename)

	sum, err := SumFile(filename)
	if err != nil {
		fmt.Errorf("Error: %w", err)
		return
	}
	fmt.Println(sum)
}
