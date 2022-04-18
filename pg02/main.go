// copy file 

package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename := os.Args[1]

	buf, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fp, err := os.Create("copy.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fp.Write(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(buf))

	// os.Stdout.Write(data)
}