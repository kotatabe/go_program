NAME = main
GO_FILES:= main.go ./...
TEXT_FILE = text/sample.txt
NUM_FILE = text/num.txt

build: 
	go build -o $(NAME) $(GO_FILES)

run-output:
	go run . -o $(TEXT_FILE)

run-copy:
	go run . -c $(TEXT_FILE)

run-sum:
	go run . -sum $(NUM_FILE)

clean:
	rm -f $(NAME)

.PHONY: build run clean run-output run-copy run-sum