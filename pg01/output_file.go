
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	
	if len(os.Args) > 2 || len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Invalid argument...")
		return
	}
	filename := os.Args[1]

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)
	}
}