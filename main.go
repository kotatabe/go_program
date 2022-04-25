package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename1 := os.Args[1]
	filename2 := os.Args[2]
	fmt.Println("==== output file =====\n")

	OutputFile(filename1)

	fmt.Println("\n\n==== copy file =====\n")
	
	CopyFile(filename1)
	
	fmt.Println("\n\n==== sum =====\n")
	sum, err := SumFile(filename2)
	if err != nil {
		fmt.Errorf("Error: %w", err)
		return
	}
	fmt.Println(sum)
}
