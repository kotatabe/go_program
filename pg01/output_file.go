
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

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(data)
}