package sumfile

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

// Create file with random num, random lines
func CreateRandFile() (*os.File, int) {
	fp, _ := os.Create("../text/rand.txt")
	defer fp.Close()

	rand.Seed(time.Now().UnixNano())
	sum := 0
	n := rand.Intn(10)
	for i := 0; i < n; i++ {
		r := rand.Intn(1000)
		fmt.Fprintln(fp, r)
		sum += r
	}
	// fp.Seek(0, 0)
	// io.Copy(os.Stdout, fp)
	return fp, sum
}

func TestSumFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{ // TODO: Add test cases.
			name:    "random numbers, random lines",
			args:    args{filename: "../text/rand.txt"},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, tt.want = CreateRandFile()
			got, err := SumFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
