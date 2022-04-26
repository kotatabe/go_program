NAME = main
GO_FILES:= main.go ./...
TEXT_FILE = text/sample.txt
NUM_FILE = text/num.txt

build: 
	go build -o $(NAME) $(GO_FILES)

run-output:
	go run . "1" $(TEXT_FILE)

run-copy:
	go run . "2" $(TEXT_FILE)

run-sum:
	go run . "3" $(NUM_FILE)

clean:
	rm -f $(NAME)

.PHONY: build run clean