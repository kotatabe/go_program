package main

import (
	"fmt"
	// "io"
	"math/rand"
	"os"
	"time"
	"testing"
)

func PutRandFile() (*os.File, int) {
	fp, _ := os.Create("text/rand.txt")
	defer fp.Close()

	rand.Seed(time.Now().UnixNano())
	sum := 0
	for_times := rand.Intn(10)
	for i := 0; i < for_times; i++ {
		r := rand.Intn(100)
		fmt.Fprintln(fp, r)
		sum += r
	}
	// fp.Seek(0, 0)
	// io.Copy(os.Stdout, fp)
	return fp, sum
}

func TestSumFile(t *testing.T) {
	_, rand_sum := PutRandFile()
	sum, e := SumFile("text/rand.txt")
	if e != nil {
		t.Errorf("%v", e)
	}
	if rand_sum != sum {
		t.Errorf("sum is %v, not %v", rand_sum, sum)
	}
}
