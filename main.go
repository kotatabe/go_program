package main

import (
	"fmt"
	// "io"
	// "math/rand"
	"github.com/kotatabe/go_program/copy_file"
	"github.com/kotatabe/go_program/output_file"
	"github.com/kotatabe/go_program/sum_file"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename := os.Args[2]
	switch os.Args[1] {
		case "1":
			fmt.Println("==== output file =====")
			output.OutputFile(filename)
		case "2":
			fmt.Println("====  copy file  =====")
			copy.CopyFile(filename)		
		case "3":
			fmt.Println("====  sum file   =====")
			sum, err := sumfile.SumFile(filename)
			if err != nil {
				fmt.Errorf("Error: %w", err)
				return
			}
			fmt.Println(sum)
		default:
			fmt.Fprintln(os.Stderr, "Invalid argument...")
			return
	}
}
