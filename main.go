package main

import (
	"flag"
	"fmt"
	"os"

	copy "github.com/kotatabe/go_program/copy_file"
	output "github.com/kotatabe/go_program/output_file"
	sumfile "github.com/kotatabe/go_program/sum_file"
	"golang.org/x/xerrors"
)

func main() {
	o := flag.Bool("o", false, "output flag")
	c := flag.Bool("c", false, "copyfile flag")
	s := flag.Bool("sum", false, "sumfile flag")
	flag.Parse()
	if len(os.Args) != 3 {
		fmt.Printf("%+v\n", xerrors.New("invalid argument"))
		return
	}
	filename := os.Args[2]

	if *o {
		fmt.Println("==== output file =====")
		output.OutputFile(filename)
	}
	if *c {
		fmt.Println("====  copy file  =====")
		copy.CopyFile(filename)
	}
	if *s {
		fmt.Println("====  sum file   =====")
		sum, err := sumfile.SumFile(filename)
		if err != nil {
			return
		}
		fmt.Println(sum)
	}
}
