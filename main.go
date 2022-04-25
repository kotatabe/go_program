package main

import (
	"fmt"
	// "io"
	// "math/rand"
	"os"
)

// func put_file() (*os.File, int) {
// 	fp, _ := os.Create("test.txt")
// 	defer fp.Close()

// 	rand.Seed(time.Now().UnixNano())
// 	sum := 0
// 	for_times := rand.Intn(10)
// 	for i := 0; i < for_times; i++ {
// 		r := rand.Intn(100)
// 		fmt.Fprintln(fp, r)
// 		sum += r
// 	}
// 	fp.Seek(0, 0)
// 	io.Copy(os.Stdout, fp)
// 	return fp, sum
// }

// func main() {
// 	put_file()
// }

func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Fprintln(os.Stderr, "Invalid argument...")
	// 	return
	// }
	filename1 := os.Args[1]
	filename2 := os.Args[2]
	fmt.Println("==== output file =====")

	OutputFile(filename1)

	fmt.Println("\n\n==== copy file =====")
	
	CopyFile(filename1)
	
	fmt.Println("\n\n==== sum =====")
	sum, err := SumFile(filename2)
	if err != nil {
		fmt.Errorf("Error: %w", err)
		return
	}
	fmt.Println(sum)
}
