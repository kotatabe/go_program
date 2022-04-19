
package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	
	if len(os.Args) > 2 || len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "Invalid argument")
		return
	}
	filename := os.Args[1]

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Open failed")
		return
	}
	defer fp.Close()

	_, err = io.Copy(io.Writer(os.Stdout), io.Reader(fp))
	// sc := bufio.NewScanner(fp)

	// for sc.Scan() {
	// 	line := sc.Text()
	// 	fmt.Println(line)
	// }
}